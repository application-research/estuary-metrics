create MATERIALIZED view content_deals_success_fail_rates
as
select ((t.success  * 1.0 /t.total  * 1.0) * 100) as "Success", ((t.failed * 1.0 / t.total * 1.0) * 100) as "Failure" from (select
    (select count(*) from content_deals as c1 where failed = false and deleted_at IS NULL) as total,
    (select count(*) from content_deals as c1 where failed = false and deal_id > 0 and deleted_at IS NULL) as success,
    (select count(*) from content_deals as c1 where failed = false and deal_id = 0 and deleted_at IS NULL) as failed) as t;