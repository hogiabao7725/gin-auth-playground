-- ENUM TYPES
CREATE TYPE user_role       AS ENUM ('user', 'organizer', 'admin');
CREATE TYPE event_status    AS ENUM ('on_sale', 'cancelled', 'ended');
CREATE TYPE booking_status  AS ENUM ('pending', 'paid', 'cancelled', 'expired');
CREATE TYPE payment_status  AS ENUM ('pending', 'succeeded', 'failed');

-- USERS
CREATE TABLE users (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(100) NOT NULL,
    email       VARCHAR(255) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,           -- bcrypt hash
    role        user_role   NOT NULL DEFAULT 'user',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- REFRESH TOKENS (for revocation support)
CREATE TABLE refresh_tokens (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash  VARCHAR(255) NOT NULL UNIQUE,    -- SHA-256 hash of raw token
    expires_at  TIMESTAMPTZ NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- EVENTS
CREATE TABLE events (
    id            UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    organizer_id  UUID        NOT NULL REFERENCES users(id),  -- user với role organizer hoặc admin
    title         VARCHAR(255) NOT NULL,
    description   TEXT,
    city          VARCHAR(100) NOT NULL,
    venue         VARCHAR(255) NOT NULL,
    banner_url    TEXT,
    start_date    TIMESTAMPTZ NOT NULL,
    end_date      TIMESTAMPTZ NOT NULL,
    status        event_status NOT NULL DEFAULT 'on_sale',
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_end_after_start CHECK (end_date > start_date)
);

-- TICKET TYPES
CREATE TABLE ticket_types (
    id               UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id         UUID    NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    name             VARCHAR(100) NOT NULL,
    description      TEXT,
    price            BIGINT  NOT NULL,
    quantity         INTEGER NOT NULL,
    sold             INTEGER NOT NULL DEFAULT 0,
    max_per_booking  INTEGER NOT NULL DEFAULT 4,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_price_non_negative   CHECK (price >= 0),
    CONSTRAINT chk_quantity_positive    CHECK (quantity > 0),
    CONSTRAINT chk_sold_valid           CHECK (sold >= 0 AND sold <= quantity),
    CONSTRAINT chk_max_per_booking_min  CHECK (max_per_booking >= 1),
    UNIQUE (event_id, name)
);

-- BOOKINGS
CREATE TABLE bookings (
    id              UUID           PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID           NOT NULL REFERENCES users(id),
    ticket_type_id  UUID           NOT NULL REFERENCES ticket_types(id),
    quantity        INTEGER        NOT NULL,
    total_amount    BIGINT         NOT NULL,
    status          booking_status NOT NULL DEFAULT 'pending',
    expires_at      TIMESTAMPTZ    NOT NULL,
    created_at      TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_booking_quantity_positive CHECK (quantity > 0),
    CONSTRAINT chk_total_amount_positive     CHECK (total_amount >= 0)
);

-- PAYMENTS
CREATE TABLE payments (
    id                  UUID           PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id          UUID           REFERENCES bookings(id) UNIQUE,
    stripe_payment_id   VARCHAR(255)   NOT NULL,   -- pi_xxx
    amount              BIGINT         NOT NULL,
    currency            VARCHAR(10)    NOT NULL DEFAULT 'vnd',
    status              payment_status NOT NULL DEFAULT 'pending',
    paid_at             TIMESTAMPTZ,
    created_at          TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);

-- INDEXES
CREATE INDEX idx_refresh_tokens_user_id          ON refresh_tokens(user_id);
CREATE INDEX idx_events_organizer_id             ON events(organizer_id);
CREATE INDEX idx_events_city                     ON events(city);
CREATE INDEX idx_events_start_date               ON events(start_date);
CREATE INDEX idx_events_status                   ON events(status);
CREATE INDEX idx_ticket_types_event_id           ON ticket_types(event_id);
CREATE INDEX idx_bookings_user_id                ON bookings(user_id);
CREATE INDEX idx_bookings_ticket_type_id         ON bookings(ticket_type_id);
CREATE INDEX idx_bookings_status                 ON bookings(status);
CREATE INDEX idx_bookings_pending_expires        ON bookings(expires_at) WHERE status = 'pending';
CREATE INDEX idx_payments_booking_id             ON payments(booking_id);
CREATE INDEX idx_payments_stripe_payment_id      ON payments(stripe_payment_id);

-- TRIGGER: auto-update updated_at
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_users_updated_at         BEFORE UPDATE ON users         FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_events_updated_at        BEFORE UPDATE ON events        FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_ticket_types_updated_at  BEFORE UPDATE ON ticket_types  FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_bookings_updated_at      BEFORE UPDATE ON bookings      FOR EACH ROW EXECUTE FUNCTION update_updated_at();
CREATE TRIGGER trg_payments_updated_at      BEFORE UPDATE ON payments      FOR EACH ROW EXECUTE FUNCTION update_updated_at();
