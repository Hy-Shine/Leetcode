/* Write your PL/SQL query statement below */
select date1 day,case when round(c2/c1,2)>0 then round(c2/c1,2) else 0 end "Cancellation Rate" 
from 
(select request_at date1,count(1) c1
from trips 
where client_id in (select users_id from users where banned='No') 
and driver_id in (select users_id from users where banned='No')
group by request_at) t
left join
(select request_at date2,count(1) c2
from trips 
where client_id in (select users_id from users where banned='No') 
and driver_id in (select users_id from users where banned='No')
and status<>'completed'
group by request_at) t2 on t.date1=t2.date2
where t.date1 between '2013-10-01' and '2013-10-03'
order by date1;