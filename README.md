# PvPtv.io (v2)

[PvPtv.io](https://pvptv.io) is a web application that helps the [World of Warcraft](https://worldofwarcraft.com/) PvP community find the best arena streamers live at any moment, for any class.

It provides a curated list of top arena players and embeds their streams directly from [Twitch](https://twitch.tv).

## Architecture (v2)

The application has been completely refactored from a client-side WASM app to a robust server-side architecture:

- **Backend**: Go 1.21+ using [Gin](https://github.com/gin-gonic/gin) for routing and HTML templates.
- **Database**: PostgreSQL for data persistence
- **Deployment**: Docker Compose with Traefik for HTTPS and routing.
- **SEO**: Fully server-rendered pages with JSON-LD structured data, sitemaps, and proper meta tags.
- **Admin**: Secure admin panel for managing channels and assignments.

## Development

### Prerequisites

- Docker & Docker Compose
- Go 1.21+ (optional, for local non-docker dev)

### Running Locally

1. **Clone the repository**
   ```bash
   git clone https://github.com/hnouts/pvptv.git
   cd pvptv
   ```

2. **Setup Environment**
   Copy the template env file:
   ```bash
   cp .template.env .env
   ```
   *Note: You'll need Twitch API credentials (`TWITCH_CLIENT_ID`, `TWITCH_CLIENT_SECRET`) to fetch live stream data.*

3. **Run with Docker Compose**
   This starts the app and a Postgres database.
   ```bash
   docker compose -f docker-compose.dev.yml up --build
   ```
   Access the site at `http://localhost:8080`.

### Admin Access

The default dev setup seeds an admin user.
- **Login**: `/admin/login`
- **Username**: `admin`
- **Password**: `admin`


## Contributing

Suggestions for new streamers are welcome! You can fill out this [form](https://forms.gle/y5wj532gvtgwWkca6).

To report issues, please open a [GitHub issue](https://github.com/hnouts/pvptv/issues).

## Support

PvPtv.io is a free tool for the community. If you like the project, you can support the development by [buying me a coffee](https://www.buymeacoffee.com/hugodev).
