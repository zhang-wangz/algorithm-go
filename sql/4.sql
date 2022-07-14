


select name from customers left join orders
on customers.id = orders.customerId where isnull(CustomerId)