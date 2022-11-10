
# Write your MySQL query statement below


select user_id, max(time_stamp) as last_stamp
from Logins
where datediff(time_stamp, "2020-1-1") >= 0 and datediff(time_stamp, "2021-1-1") < 0
group by user_id;

# Write your MySQL query statement below


select user_id, max(time_stamp) as last_stamp
from Logins
where year(time_stamp) = 2020
group by user_id;