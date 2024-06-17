package com.bocabbage.newssubscriber.monodemo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cache.annotation.EnableCaching;

@SpringBootApplication
@EnableCaching
public class NewsSubscriberMonoDemoApplication {
	public static void main(String[] args) {
		SpringApplication.run(NewsSubscriberMonoDemoApplication.class, args);
	}
}
