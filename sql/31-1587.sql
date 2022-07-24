
# Write your MySQL query statement below

select u.name as NAME, sum(
    amount
) as BALANCE
from Users u, Transactions t 
where u.account = t.account
group by u.account
having sum(amount) > 10000