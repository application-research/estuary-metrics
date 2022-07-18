-- trigger event upon insert on content_deals
CREATE TRIGGER content_deals
    AFTER INSERT OR UPDATE OR DELETE ON content_deals
    FOR EACH ROW EXECUTE PROCEDURE notify_event();