## Go advanced bootcamp week 8 homework

## Problems

1. Use the redis benchmark tool | to test the performance of redis get set for 10 20 50 100 200 1k 5k bytes value size.

2. Write a certain amount of kv data | Evaluate yourself based on the data size 1w-50w | Combine the info memory information before and after writing | Analyse the average memory space occupied by each key under the different value sizes mentioned above.

## Environment Preparation

1. Installation and configuration of Redis - [Click for details](. /redis_install.md)

2. Monitor using Grafana + Prometheus + redis_exporter - [Click for details](. /monitor.md)

## benchmark solution

1. using Centos 7.9 - 2 cores 4GB virtual machine for testing

2. test for Jemalloc and Libc memory allocation respectively.

3. use Redis' own benchmark tool for testing

4. test 10 20 50 100 200 1k 5k bytes value size, each 50w times, key randomly generated 50w times

5. [Test command details](. /benchmark-command.md)

## redis get set benchmark performance

Requests|ValueSize|Set req/sec(jemalloc)|Get req/sec(jemalloc)|Set req/sec(libc)|Get req/sec(libc)
-|-|-|-|-|-|-
50w|10|127583.57|126167.05|99840.26|98096.91
50w|20|117980.18|122639.19|137854.97|125219.13
50w|50|98541.59|137400.39|92267.95|97068.53
50w|100|140173.81|139586.83|106837.61|96786.69
50w|200|109793.59|94500.09|115500.12|99108.02
50w|500|103971.72|97924.02|127616.13|136276.92
50w|1024|95803.8|95328.89|109433.14|106360.34
50w|5120|65971.77|92387.28|83194.67|91224.23

! [Memory usage graph](. /img/redis_req_sec.png)

### Observing the above test results, the following points are obtained

- jemalloc efficiency is inversely proportional to value size.
- The efficiency of libc shows a jittery trend.
- value size in 10~100 jemalloc way set, get are more efficient than libc.
- value size in 200~5120 libc way set, get are more efficient than jemalloc (at 5120 jemalloc get efficiency is closer to libc)

## Write memory analysis - test results

>> 50w requests for randomly generated keys, the final number of keys fluctuates between 393000~394000 (randomly generated keys will be duplicated, some of the keys will be repeatedly updated), after which the average value is calculated using the median of 393500.

Requests|ValueSize|memory(jemalloc)|avg(jemalloc)|memory(libc)|avg(libc)
-|-|-|-|-|-|-
50w|10|34.8M|93|37.9M|101
50w|20|37.9M|101|37.9M|101
50w|50|49.9M|133|52.9M|141
50w|100|70.9M|189|70.9M|189
50w|200|113M|302|113M|302
50w|500|221M|589|227M|605
50w|1024|509M|1357|419M|1117
50w|5120|2.28G|6222|1.91G|5212

! [Memory usage graph](. /img/redis_memory_avg.png)

### Observing the above test results, the following points are obtained

1. value size between 10 ~ 500 byte jemalloc memory allocator occupies less memory than libc;

2. jemalloc memory allocator occupies more memory than libc when the value size is 1024 byte and above. 3;

3. the above two articles can not be concluded that the value size size of jemalloc and libc allocation of memory there is a standard threshold, but also need to start from the principle of the two allocation algorithms in order to start the comparison of screening;

4. from the above average memory usage, we can see that there is a fixed size difference between the actual memory usage and the value size.

   - From the help of benchmark, we can see that the key is a 12-bit number, so the key needs to take up 8 bytes.

   - By learning the internal data structure of redis, we know that when saving key or value, we need to use RedsiObject structure to save it in redis, which takes up 16 byte+string size memory for a string.

   - The hash table that holds the key itself uses 24 bytes of memory, and due to the memory allocator it actually takes up 32 bytes

   - (key = 8 + 8) + (value = 16 + valueSize) + (hash entry = 32) = 64 + valueSize + the extra memory allocated by the memory allocator due to multiplication.

## Monitor

### grafana - jemalloc

! [jemalloc memory usage graph](. /img/redis_memory_jemalloc.png)

### grafana - libc

! [libc memory usage graph](. /img/redis_memory_libc.png)

Translated with DeepL.com (free version)
