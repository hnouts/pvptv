Totally hear you — we’ve got enough to finalize the refactor. Here’s a tight, hand-offable spec that keeps the exact public UI/CSS/UX and swaps the internals.

# PvPtv Refactor Spec (v1)

## Goals

- Security + maintainability; data agility via DB + admin.
- Public site: **unchanged markup/CSS and navigation**.
- Remove **Plunderstorm** (+ 301 from `/plunderstorm*` → `/`).

## Stack

- **Go** ≥ 1.21, **Gin** + `html/template` (+ HTMX sprinkles), no WASM.
- **Postgres** (VPS), access via **sqlc**.
- **Docker Compose** + **Traefik** (HTTPS).
- **Self-hosted Plausible** (`plausible.hnts.dev`) tracking `pvptv.hnts.dev`.
- **In-memory LRU cache** (Helix TTL = **60s**).
- Logs: JSON to stdout; `/healthz` & `/readyz`.

## Data model (Postgres, normalized)

```sql
-- games optional later; for now classes.game_slug='wow'
CREATE TABLE classes (
  id           BIGSERIAL PRIMARY KEY,
  game_slug    TEXT NOT NULL,               -- 'wow' (room for 'lol' later)
  class_slug   TEXT NOT NULL,
  class_name   TEXT NOT NULL,
  UNIQUE (game_slug, class_slug)
);

CREATE TABLE specs (
  id         BIGSERIAL PRIMARY KEY,
  class_id   BIGINT NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  spec_slug  TEXT NOT NULL,
  spec_name  TEXT NOT NULL,
  UNIQUE (class_id, spec_slug)
);

CREATE TABLE channels (
  id                 UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  twitch_login       TEXT NOT NULL UNIQUE,     -- canonical key
  twitch_user_id     TEXT UNIQUE,              -- filled lazily
  display_name       TEXT NOT NULL,
  language           TEXT,                     -- e.g. 'en'
  region             TEXT,                     -- free text for now
  mature_flag        BOOLEAN DEFAULT FALSE,
  is_published       BOOLEAN DEFAULT TRUE,
  sort_weight        INT DEFAULT 0,
  notes              TEXT,
  avatar_url         TEXT,
  bio                TEXT,
  auto_sync_avatar   BOOLEAN DEFAULT FALSE,
  auto_sync_bio      BOOLEAN DEFAULT FALSE,
  socials            JSONB DEFAULT '{}'::jsonb,
  created_at         TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at         TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE channel_specs (
  channel_id  UUID NOT NULL REFERENCES channels(id) ON DELETE CASCADE,
  spec_id     BIGINT NOT NULL REFERENCES specs(id) ON DELETE CASCADE,
  is_primary  BOOLEAN NOT NULL DEFAULT FALSE,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
  PRIMARY KEY (channel_id, spec_id)
);
-- enforce at most one primary per channel
CREATE UNIQUE INDEX channel_specs_one_primary
  ON channel_specs(channel_id) WHERE is_primary = TRUE;
```

## Public site (unchanged UX)

- **Keep existing layout & CSS**; reimplement pages server-rendered with identical markup.
- **Sorting:** Online first → viewer count (desc) → name (A–Z).
- **Twitch embed**: `parent` whitelisted for `pvptv.hnts.dev`, `localhost`, `127.0.0.1`.
- **Live data:** On request, call Helix `Get Streams` (batch by `user_login`, up to 100), cache per-login & list for **60s**. Auto-refresh token via client-credentials.

## Admin (/admin)

- **Auth:** single seeded admin, Argon2id hash; HTTP-only `SameSite=Lax` cookie; **7-day** session, idle **12h**; optional “remember me” (30d). CSRF via double-submit token.
- **Pages:**

  - Channels list (search by name/login; filter by class/spec/language; bulk publish/unpublish; inline sort_weight).
  - Create/Edit channel (multi class/spec, one primary; “Test Twitch login”; optional “Fetch from Twitch” to fill display_name/user_id).
  - Live preview of public card.

- **Import tool (one-time):** parse current Go catalog, **drop Plunderstorm**, insert channels/classes/specs, map multi-spec + primary.

## Security & headers

- HSTS (includeSubDomains; preload), Referrer-Policy: `strict-origin-when-cross-origin`, basic `Permissions-Policy` (camera/mic/geolocation=“()”).
- Rate-limit login (e.g., 5 attempts/15 min/IP).
- No CSP for now (per choice).

## Config / env & secrets

- `TWITCH_CLIENT_ID` (secret), `TWITCH_CLIENT_SECRET` (secret)
- `DATABASE_URL` (secret), `SESSION_KEY` (secret)
- `ADMIN_USERNAME`, `ADMIN_PASSWORD_HASH`
- `PLAUSIBLE_HOST=plausible.hnts.dev`, `PLAUSIBLE_DOMAIN=pvptv.hnts.dev`
- All injected via **Docker secrets/env_file**; never committed.

## Docker/Traefik (sketch)

```yaml
services:
  app:
    image: pvptv/app:latest
    env_file: .env
    secrets: [db_url, session_key, twitch_client_id, twitch_client_secret]
    depends_on: [db]
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.pvptv.rule=Host(`pvptv.hnts.dev`)"
      - "traefik.http.routers.pvptv.entrypoints=websecure"
      - "traefik.http.routers.pvptv.tls.certresolver=lets"
  db:
    image: postgres:16
    environment:
      POSTGRES_DB: pvptv
      POSTGRES_USER: pvptv
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
secrets:
  db_url: { file: ./secrets/DATABASE_URL }
  session_key: { file: ./secrets/SESSION_KEY }
  twitch_client_id: { file: ./secrets/TWITCH_CLIENT_ID }
  twitch_client_secret: { file: ./secrets/TWITCH_CLIENT_SECRET }
```

## Migrations & codegen

- **golang-migrate** SQL files in `/migrations`.
- **sqlc** for typed queries in `/internal/db`.
- Repo layout:

```
/cmd/server        # main
/internal/http     # handlers, middleware, templates
/internal/twitch   # helix client + cache
/internal/admin    # admin handlers/forms
/internal/db       # sqlc gen
/migrations        # SQL
/assets            # CSS/JS/images (public CSS copied 1:1)
/templates         # public + admin templates (public must match current markup)
```

## SEO

- **General approach**

  - Server-render all public pages with stable, crawlable URLs and minimal client-side navigation (HTMX used progressively, not for critical routing).
  - Keep existing DOM/CSS for visible content; **all SEO work is done via `<head>` tags, HTTP headers, and structured data**.

- **HTML `<title>` and meta descriptions**

  - Define a consistent title pattern:
    - Home: `PvPtv – High-level WoW PvP stream hub`
    - Class/spec pages (or filtered views): `ClassName SpecName – WoW PvP streams on PvPtv`
    - Channel detail: `DisplayName – WoW PvP streamer on PvPtv`
    - Generic fallback: `PvPtv – WoW PvP streams`
  - Render a unique `<meta name="description">` per key view:
    - Home: short pitch of the site and value.
    - Class/spec listing: what’s special about that spec’s PvP streams.
    - Channel detail: who the streamer is, language, region, class/spec.
  - Enforce max ~150–160 chars for descriptions, 55–60 chars for critical titles.

- **Canonical URLs & redirects**

  - Each crawlable page emits `<link rel="canonical" href="...">` with the preferred `https://pvptv.hnts.dev/...` URL.
  - Normalize:
    - Redirect `http` → `https`.
    - Redirect `www.pvptv.hnts.dev` → `pvptv.hnts.dev` (single canonical host).
    - Redirect trailing slashes consistently (e.g. `'/c/foo/' → '/c/foo'`).
    - Keep the spec’s 301s: `/plunderstorm*` → `/`.
  - Avoid duplicate content paths for the same logical page.

- **Robots & sitemap**

  - Serve `robots.txt` at `/robots.txt`:
    - Allow crawl for public pages; disallow `/admin` and non-public APIs.
    - Include `Sitemap: https://pvptv.hnts.dev/sitemap.xml`.
  - Serve `sitemap.xml`:
    - At minimum: home, any class/spec listing pages, each public channel page `/c/:twitch_login`.
    - Include `lastmod` timestamps from DB (`channels.updated_at`) where available.
    - Regenerate on relevant DB changes or nightly.

- **Open Graph / social cards**

  - For all public pages, include:
    - `og:type` (`website` for home/list pages, `profile` or `article`-like for channel detail).
    - `og:title`, `og:description`, `og:url` (matching canonical), and `og:image`.
  - Channel detail pages:
    - Prefer channel avatar if available; fallback to a generic PvPtv image.
    - Include `twitter:card=summary_large_image`, `twitter:title`, `twitter:description`, `twitter:image`.
  - Reuse the same content as `<title>`/description where possible to avoid drift.

- **Structured data (JSON-LD)**

  - Add JSON-LD `<script type="application/ld+json">` in `<head>`; no layout changes.
  - Recommended types:
    - **Site-wide (home page)**: `WebSite` with a `SearchAction` (if you have site search).
    - **Listing views (e.g. main stream grid)**: `CollectionPage` or `ItemList` of streamers; each item links to the channel’s PvPtv page.
    - **Channel detail**:
      - Use `Person` or `Organization` (depending on branding) with:
        - `name`, `url` (PvPtv page), `sameAs` (Twitch and other socials from `socials` JSONB), `image` (avatar).
      - Optionally, an `AboutPage` or `ProfilePage` referencing that entity.
  - Ensure JSON-LD is valid and consistent (no broken `@id` URLs; test via Google Rich Results / schema.org validators).

- **Content & internal linking (within existing layout)**

  - Without changing CSS/DOM classes, allow for **small, text-based enhancements** in existing containers:
    - Short, human-readable copy on the home page explaining what PvPtv is and who it’s for.
    - One or two lines on class/spec views describing that spec in PvP PvE/PvP context (no keyword stuffing).
  - Ensure internal links:
    - From the home/listing page to each channel `/c/:twitch_login`.
    - Back-links from channel pages to their class/spec category views.
    - Use descriptive link text (e.g., class/spec + streamer name) rather than “click here”.

- **Performance & Core Web Vitals**

  - Leverage existing server-side rendering and add:
    - Gzip or Brotli for HTML/CSS/JS via Traefik/app config.
    - Far-future cache headers for static assets under `/assets` (fingerprinted filenames).
    - Reasonable `Cache-Control` for HTML (e.g., low TTL + ETag/Last-Modified).
  - Optimize:
    - Critical CSS is already static; ensure CSS/JS payload size is small and shipped once.
    - Lazy-load non-critical images where feasible without changing visible layout.
  - Monitor Core Web Vitals (Largest Contentful Paint, Cumulative Layout Shift, Interaction to Next Paint) via Plausible or an external tool.

- **Internationalization / language**

  - Set `<html lang="en">` (or correct primary language) on all pages.
  - If/when multiple languages are supported, add `hreflang` link tags and make sure URLs are structured to reflect language (e.g., `/en/...`, `/fr/...`), but that’s out of scope for v2.

- **Crawling, indexing, and health**

  - Expose a **crawl health** check (can be same as `/readyz`) verifying:
    - `robots.txt` is served.
    - `sitemap.xml` is valid.
    - Root page returns `200` with non-empty HTML.
  - Avoid:
    - `noindex` on public pages (reserve `noindex`/`nofollow` for `/admin` and system pages).
    - Fragment identifiers (`#!`) or JS-only navigation for primary content.

- **Acceptance criteria (SEO)**
  - All public pages have unique, descriptive `<title>`, meta description, canonical URL, and Open Graph/Twitter metadata.
  - `robots.txt` and `sitemap.xml` are reachable and list all public channel pages.
  - Channel detail pages render valid `Person` (or equivalent) JSON-LD with Twitch profile in `sameAs`.
  - Home and major listing pages render valid `WebSite`/`CollectionPage` JSON-LD.
  - Core Web Vitals scores are “Good” for desktop and mobile on synthetic tests for the home page and a representative channel page.

## Redirects & routes

- 301: `/plunderstorm` and `/plunderstorm/*` → `/`.
- Public routes mirror current pages; channel pages at `/c/:twitch_login` (fine to change internally as long as UI is same).
- Admin routes: `/admin/login`, `/admin/logout`, `/admin/channels`, `/admin/channels/new`, `/admin/channels/:id`.

## Acceptance criteria

- Pixel-parity for public pages (same DOM/CSS classes).
- Sorting & live/idle grouping identical to today.
- Admin can CRUD channels, assign multiple specs w/ single primary, bulk publish/unpublish, search/filter, and preview.
- Plunderstorm fully removed + redirects in place.
- No secrets in repo; Twitch embeds work on prod & dev.
