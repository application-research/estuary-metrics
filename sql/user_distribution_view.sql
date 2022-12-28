create MATERIALIZED VIEW user_distribution
AS
select a.username, a.id,sum(c.size) from users as a, content_deals as b, contents as c
where a.id = b.user_id and b.id = c.id and b.user_id = a.id group by a.id order by sum(c.size) desc;

