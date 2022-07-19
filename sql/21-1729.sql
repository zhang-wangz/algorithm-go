

# Write your MySQL query statement below


select user_id, count(follower_id)  followers_count
from Followers
group by user_id
order by user_id