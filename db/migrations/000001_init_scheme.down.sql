-- Drop triggers
DROP TRIGGER IF EXISTS trg_payments_updated_at ON payments;
DROP TRIGGER IF EXISTS trg_bookings_updated_at ON bookings;
DROP TRIGGER IF EXISTS trg_ticket_types_updated_at ON ticket_types;
DROP TRIGGER IF EXISTS trg_events_updated_at ON events;
DROP TRIGGER IF EXISTS trg_users_updated_at ON users;

-- Drop function
DROP FUNCTION IF EXISTS update_updated_at();

-- Drop tables
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS bookings;
DROP TABLE IF EXISTS ticket_types;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS users;

-- Drop ENUM types
DROP TYPE IF EXISTS payment_status;
DROP TYPE IF EXISTS booking_status;
DROP TYPE IF EXISTS event_status;
DROP TYPE IF EXISTS user_role;
