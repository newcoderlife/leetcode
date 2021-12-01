select t.Email
from (select count(*) as num, Email
      from Person
      group by Email) t
where t.num > 1