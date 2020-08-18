Adnymics

Hi Sir,
Thanks for giving me this opportunty. The tasks are tricky but interesting as well and I learned some new things from these tasks. Personally I my favoraite is task2. 
I have created seprate folder for different tasks. Following are the solutions and explations as you have requested.

Task 1.a:
As per the structure of the question I can see that task is tricky but the solution will most probably be simple.
I am not able to fully understand the requirements, so I decided not to spend much time on this task. 
Please let me know if you can provide some more details on this task, I will be happy to complete this.

Task 1.b:

find / -printf '%s %p\n'| sort -nr | head -3

This task is pretty simple.
find - is a very useful and handy command to search for files from the command line. 
/ - is the path of dir.
printf - printf prints a formatted string to the standard output. Its roots are in the C programming language, which uses a function by the same name. %s is a string and %p is a memory address pointer.
sort - sort is a simple and very useful command which will rearrange the lines in a text file so that they are sorted, numerically and alphabetically. -n, --numeric-sort	Compare according to string numerical value. -r, --reverse	Reverse the result of comparisons.
head, by default, prints the first 10 lines of each FILE to standard output.

Task 2
As I haven't worked on keepalived and HAProxy before so it took me alot of time to do this task. As I ran out of time, I decided not to setup HAProxy. If you can provide more time, I am sure that I can complete this task.

To run the setup first run image builder script.
    
    chmod +x imageBuilder.sh
    ./imageBuilder.sh

This script will create the master and backup docker images for you.
After this make sure that the subnet 172.18.0.0/16 is available otherwise the setup won't work.
Now run docker compose file.

    docker-compose up

You will see master running at port 8880 and port 8881 will be unreachable. Now try to pause  keepalived_master.

     docker pause keepalived_master

You will see that port 8880 is now unavailable and backup is running on port 8881.

Task 3

LCU of new connections: 26.6/25 = 1.064 LCU
Active connections in a min: 27*60 = 1620
LCU of active connections: 1620/3000 = 0.54 LCU
LCU of processed Bytes: 0.9/1 = 0.9 LCU

Out of 3 dimenshions we will choose the highest LCU which is 1.064
LCU charges per hour: 1.064*$0.008 = $0.008512

To calculate the monthly chrges we need to add the alb hourly charges as well.
LCU hourl charges + Alb hourly charge: $0.008512+$0.0225 = $0.031012
Monthly charge: $0.031012*24(hours)*30(days) = $22.32

I even used AWS Calculator and it showed the cost of $26.

Task 4

I have attached a ec2.yml file. This is a cloudformation template to create all the resources that you have requested.

Task 5
How would you scale a monolithic API service that is running in the cloud which fails to process API calls fast enough?

In real life scenario we need to do rca and figure the exact reason why the API calls are not fast.

If there are not a large number of requests but each request requires lot of resorces then we can use Verticle scaling. We can increase the number of CPU and Memory available to the instance such that the time to process each request will decrease.

If a single request doesn't require a lot of resources but there are large number of request being sent then we can use Horizontal scaling. We can put the instance in an autoscaling group and use scaling policies based on cpu usage or number of requests.

In long term we can collaborate with dev team to increase the performance of the application or use rate limiting or convert it to a micro service such that we can deploy it in a container based environment like kubernetes.

Task 6

First of all I don't code in python and I hope choice of language is not an issue. I have written the solution in golang as I currenlt work on shell and golang only.
I have assumed "bitcoin" as case sensitive and created solution based on this. The solution is in file bitcoin.go, you will require golang 1.13 or above to run this program. I have taken data from the file and created a smaller sample data in test.txt 
To run this program you can simply type - go run bitcoin.go

Just a suggestion, the files provided by wikipedia requires a lot of resources to process. A lot time got wasted in opening and searching inside the file. People who don't have high performance laptops will find this task more difficult.
Rather than giving link to wikipedia file, please create a sample data file and attach it with the mail.

You can drop me a mail in case of any issues.

Thanks,
Keshav