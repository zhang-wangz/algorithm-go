


select firstName,lastName,city, state
from Person p1
         left join Address p2 on p1.personId = p2.personId