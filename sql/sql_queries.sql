select * from content_deals
where failed = false;
select count(*) from obj_refs;

select * from obj_refs a, content_deals b
where a.content = b.content; // 36

select * from obj_refs a, content_deals b, contents c
where a.content = b.content and b.content = c.id;


select ((t.failed * 1.0 / t.total * 1.0) * 100) as "fail_rate",((t.success  * 1.0 /t.total  * 1.0) * 100) as "success_rate" from (select
                                                                                                                                      (select count(*) from content_deals as c1 where deal_id > 0) as total,

select count(*) from content_deals as c1, contents as c2
where c1.id = c2.id; // 813495 (total)

select count(*) from content_deals
where deal_id > 0; // 125286 (content with deals)

select count(*) from content_deals
where deal_id = 0; // 688209 (content without deals)

select count(*) from content_deals
where failed = false and deal_id > 0; // 106251 (accepted but with deal)

select count(*) from content_deals
where failed = false and deal_id = 0; // 19397 (accepted but no deal)

select (count(*)) from content_deals
where failed = true and deal_id > 0;  // 19035 (failed but with deal)

select count(*) from content_deals
where failed = true and deal_id = 0; // 668812 (failed but no deal)


select count(*) from contents
where failed = true; // 1946805

select count(*) from contents
where failed = false; // 27053339


select count(*) from content_deals
where deal_id != 0; // 124547

select count(*) from content_deals
where deal_id = 0; // 680920

///
select * from content_deals;
select count(*) from content_deals
where user_id != 0;

//
select count(*) from retrieval_success_records;
select count(*) from retrieval_failure_records;

//  distribution of data size uploaded per user
select a.username, a.id,sum(c.size) from users as a, content_deals as b, contents as c
where a.id = b.user_id and b.id = c.id and b.user_id = a.id group by a.id order by sum(c.size) desc;

//   time to a successful deal (from content)
select a.created_at, b.transfer_finished from contents as a, content_deals as b
where a.id = b.id and b.deal_id != 0;

//   time to a successful deal (from content deals)
select a.created_at, b.transfer_finished from contents as a, content_deals as b
where a.id = b.id and b.deal_id != 0;


// total retrieval

(select
    (select count(*) from retrieval_failure_records) as retrieval_failure_records,
    (select count(*) from retrieval_success_records) as retrieval_success_records) as total


(select
    (select count(*) from retrieval_failure_records) as retrieval_failure_records, // 171
    (select count(*) from retrieval_success_records) as retrieval_success_records) // 25


select * from retrieval_success_records

    //  uundialable miners
select count(*) from retrieval_failure_records as a where a.phase = 'query'



    // collections spread
select a.user_id, count(*) from collections a, users b where a.user_id = b.id group by a.user_id order by count(*) desc limit 10;

select count(*) from invite_codes a where a.claimed_by is not null;
select count(*) from invite_codes a where a.claimed_by = 0;
select count(*) from invite_codes a where a.claimed_by = 0;


select count(*) from storage_miners where suspended = false;
select count(*) from storage_miners where suspended = true;


select count(*) from auth_tokens;

select count(*) from auth_tokens a where a.deleted_at is not null;
select count(*) from auth_tokens a where a.deleted_at is null;
select count(*) from auth_tokens a where a.expiry > now(); // no expired
select count(*) from auth_tokens a where a.expiry < now(); // expired


select count(*) from content_deals a where a.on_chain_at is not null;
select count(*) from content_deals a where a.on_chain_at is null;

select count(*) from content_deals a where a.transfer_started is not null and transfer_finished is not null;
select count(*) from content_deals a where a.verified;
select count(*) from content_deals a where a.verified = false;

select a.miner, count(*) from content_deals a group by a.miner order by count(*) desc;
select a.miner, count(*) from content_deals a group by a.miner order by count(*) desc limit 10;

select a.username, a.id,sum(c.size) from users as a, content_deals as b, contents as c
where a.id = b.user_id and b.id = c.id and b.user_id = a.id group by a.id order by sum(c.size) desc limit 10;

select a.miner, b.name, count(*) from content_deals a, storage_miners b where a.miner = b.address group by a.miner, b.name order by count(*) desc limit 10;

select a.miner,  b.name, sum(a.size) from retrieval_success_records a, storage_miners b where a.miner = b.address  group by a.miner,b.name order by sum(a.size) desc limit 10;

select a.miner, b.name, sum(cast(total_payment as numeric)) from retrieval_success_records a, storage_miners b where a.miner = b.address group by a.miner,b.name order by sum(cast(total_payment as numeric)) desc limit 10;