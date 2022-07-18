

# Write your MySQL query statement below

select id,
       case when p_id is null then 'Root'
            when id in (select distinct p_id from tree) then 'Inner' else 'Leaf' end as type
from tree
order by id