package com.example.cachedemo;

import com.example.cachedemo.model.Event;
import java.sql.Timestamp;
import java.time.LocalDateTime;
import java.time.temporal.Temporal;
import java.util.HashMap;
import java.util.Map;
import java.util.PriorityQueue;

public class MyCache {
    /*
    *  Each key should have a TTL
    *  Memory is important ==> We cannot have stale entries in the cache
    * */

    private static final Map<String, Event> cache = new HashMap<>();
    private static final PriorityQueue<Event> pq = new PriorityQueue<>((a, b) ->
    {
        return (int) ( (a.getCreationTime()+a.getTtl()) - (b.getCreationTime() + b.getTtl()) );
    });

    //4:53 ttl 5 mins
    // 4:55 ttl 2 mins
    public void addEvent(String key, String value, Long ttl){
        Event event = Event.builder().key(key).value(value).creationTime(
                System.currentTimeMillis()).ttl(ttl).build();
        //This part can be made asynchronous
        while(pq.size() > 0 & pq.peek().getTtl() + pq.peek().getCreationTime() < System.currentTimeMillis()){
            Event expiredEvent = pq.poll();
            cache.remove(expiredEvent.getKey());
        }


        cache.put(key, event);
    }

    public String getEvent(String key){

        Event event = cache.get(key);
        if(event == null){
            return "Key not found in cache";
        }

        return event.getValue();

    }


}
