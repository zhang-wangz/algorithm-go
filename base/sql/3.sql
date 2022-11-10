# Write your MySQL query statement below

# 与null值作比较需要单独列出来
select name from customer where referee_id != 2  or referee_id is Null