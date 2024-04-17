package com.example.cachedemo.model;

import java.sql.Timestamp;
import java.time.LocalDateTime;
import java.time.temporal.Temporal;
import java.time.temporal.TemporalUnit;
import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

@Builder
@Getter
@Setter
public class Event {
    private String key;
    private String value;
    private Long creationTime;
    private Long ttl;


}
