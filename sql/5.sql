
select employee_id, if(employee_id % 2 =1 and left(name,1)!='M', salary, 0) as bonus
from Employees
order by employee_id;

# case when then else end
select employee_id,
       case
           when employee_id%2=1 and left(name,1)!='M' then salary else 0
           end as bonus
from Employees
order by employee_id

