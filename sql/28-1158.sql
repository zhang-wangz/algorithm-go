
# Write your MySQL query statement below


select user_id as buyer_id, join_date,
       (case when order_id is null
                 then 0 else count(*) end) as orders_in_2019
from Users u
         left join Orders o on u.user_id = o.buyer_id
    and year(order_date) = 2019
group by user_id
