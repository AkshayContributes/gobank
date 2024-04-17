package com.aksact.ratelimiterdemo.controller;

import io.github.bucket4j.Bandwidth;
import io.github.bucket4j.Bucket;
import io.github.bucket4j.Bucket4j;
import io.github.bucket4j.Refill;
import java.time.Duration;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
@RequestMapping("/rate-limit")
public class RateLimitController {

    @GetMapping("/token")
    public ResponseEntity<String> getToken() {
        //Refill
        Refill refill = Refill.of(10, Duration.ofSeconds(1));

        Bucket bucket = Bucket4j.builder()
                .addLimit(Bandwidth.classic(5, refill))
                .build();

        return ResponseEntity.ok("Token Generated Successfully");
    }

    @GetMapping("/demo")



}
