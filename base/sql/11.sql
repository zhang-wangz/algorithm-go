
# Write your MySQL query statement below

select employee_id
from
    (
        select employee_id from Employees
        union all
        select employee_id from Salaries
    ) as t
group by employee_id
having count(employee_id) = 1
order by employee_id