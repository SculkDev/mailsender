# mailsender

HTTP service that renders email templates and sends them via SMTP. Errors are reported to Sentry.

## API

`POST /mailhook`

Requires `X-Internal-Key` header (must match `MAILSENDER_INTERNAL_KEY` env var).

Request body (JSON):

```json
{
  "template_type": "recovery_valid",
  "to": "user@example.com",
  "recovery_url": "https://example.com/recover?token=abc"
}
```

The service looks for template files in `./templates/<template_type>/` directory. Each template type needs three files: `html.template`, `text.template`, `subject.template`.

Variables in templates use `{{ variable_name }}` syntax.

Some template types are ignored (return 200 without sending): `recovery_invalid`, `recovery_code_invalid`, `verification_invalid`, `verification_code_invalid`.

## Available request fields

`body`, `login_code`, `recipient`, `recovery_code`, `recovery_url`, `registration_code`, `subject`, `template_type`, `to`, `verification_code`, `verification_url`

All fields are optional except `template_type` and `to` (or `recipient`).

If `subject` is provided in the request, it overrides the rendered subject template.

## Environment variables

See `.env.example`. Required:

- `SMTP_HOST`, `SMTP_PORT`, `SMTP_USERNAME`, `SMTP_PASSWORD` — SMTP server
- `SMTP_FROM` — sender address
- `MAILSENDER_INTERNAL_KEY` — auth key for incoming requests
- `SENTRY_DSN` — Sentry DSN (optional)

## Run

```
go build -o mailsender .
./mailsender
```

Listens on `:1432`.

## Docker

```
docker compose up --build
```
