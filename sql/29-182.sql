# Write your MySQL query statement below

select Email from  Person
group by Email 
having count(Email) > 1