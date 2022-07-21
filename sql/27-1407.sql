

# Write your MySQL query statement below

select u.name, if(sum(r.distance) is NUll, 0, sum(r.distance)) as travelled_distance
from Users u
         left join Rides r on r.user_id = u.id
group by user_id
order by travelled_distance desc, name asc