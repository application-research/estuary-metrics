create MATERIALIZED view total_objects_size
AS
select sum(size) from objects where cid in (select cid from content_deals where cid is not null and deal_id > 0 and failed = false);


create MATERIALIZED view total_objects_and_objects_ref
AS
select (
        (select count(*) from obj_refs as total_obj_refs),
        (select count(*) from objects as total_objects)
           );