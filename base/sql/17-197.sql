
# Write your MySQL query statement below

select Weather.id from Weather
                           left join Weather w on DATEDIFF(Weather.recordDate, w.recordDate) = 1
where Weather.Temperature > w.Temperature
