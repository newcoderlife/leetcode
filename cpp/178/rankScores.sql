select
  a.score as Score,
  count(distinct b.score) as Rank
from
  scores a,
  scores b
where
  b.score >= a.score
group by
  a.id
order by
  a.score DESC;