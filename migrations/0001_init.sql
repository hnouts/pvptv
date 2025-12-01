-- Initial PvPtv schema (v2)

-- Enable UUID generation for channels.id
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- games optional later; for now classes.game_slug = 'wow'
CREATE TABLE classes (
  id         BIGSERIAL PRIMARY KEY,
  game_slug  TEXT NOT NULL,               -- 'wow' (room for 'lol' later)
  class_slug TEXT NOT NULL,
  class_name TEXT NOT NULL,
  UNIQUE (game_slug, class_slug)
);

CREATE TABLE specs (
  id        BIGSERIAL PRIMARY KEY,
  class_id  BIGINT NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
  spec_slug TEXT NOT NULL,
  spec_name TEXT NOT NULL,
  UNIQUE (class_id, spec_slug)
);

CREATE TABLE channels (
  id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  twitch_login     TEXT NOT NULL UNIQUE,     -- canonical key
  twitch_user_id   TEXT UNIQUE,              -- filled lazily
  display_name     TEXT NOT NULL,
  language         TEXT,                     -- e.g. 'en'
  region           TEXT,                     -- free text for now
  mature_flag      BOOLEAN DEFAULT FALSE,
  is_published     BOOLEAN DEFAULT TRUE,
  sort_weight      INT DEFAULT 0,
  notes            TEXT,
  avatar_url       TEXT,
  bio              TEXT,
  auto_sync_avatar BOOLEAN DEFAULT FALSE,
  auto_sync_bio    BOOLEAN DEFAULT FALSE,
  socials          JSONB DEFAULT '{}'::jsonb,
  created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE channel_specs (
  channel_id UUID NOT NULL REFERENCES channels(id) ON DELETE CASCADE,
  spec_id    BIGINT NOT NULL REFERENCES specs(id) ON DELETE CASCADE,
  is_primary BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  PRIMARY KEY (channel_id, spec_id)
);

-- enforce at most one primary per channel
CREATE UNIQUE INDEX channel_specs_one_primary
  ON channel_specs(channel_id) WHERE is_primary = TRUE;


