package com.bocabbage.newssubscriber.monodemo.web.configurations;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cache.CacheManager;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.cache.RedisCacheConfiguration;
import org.springframework.data.redis.cache.RedisCacheManager;
import org.springframework.data.redis.connection.RedisConnectionFactory;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.serializer.*;


import java.time.Duration;
import java.util.HashMap;
import java.util.Map;

@Configuration
public class RedisConfig {

    // Serialization related
    @Autowired
    private GenericJackson2JsonRedisSerializer genericJackson2JsonRedisSerializer;

    @Autowired
    private StringRedisSerializer stringRedisSerializer;

    @Bean
    public RedisTemplate<String, Object> redisTemplate(RedisConnectionFactory redisConnectionFactory) {
        RedisTemplate<String, Object> template = new RedisTemplate<>();
        template.setConnectionFactory(redisConnectionFactory);

        template.setKeySerializer(stringRedisSerializer);
        template.setValueSerializer(genericJackson2JsonRedisSerializer);
        template.setHashKeySerializer(stringRedisSerializer);
        template.setHashValueSerializer(genericJackson2JsonRedisSerializer);

        template.afterPropertiesSet();
        return template;
    }

    // Cache config
    @Bean
    public CacheManager cacheManager(RedisConnectionFactory redisConnectionFactory) {
        RedisCacheConfiguration defaultCacheConfig = RedisCacheConfiguration.defaultCacheConfig()
                // Use jackson serializer
                .serializeKeysWith(RedisSerializationContext.SerializationPair.fromSerializer(stringRedisSerializer))
                .serializeValuesWith(RedisSerializationContext.SerializationPair.fromSerializer(genericJackson2JsonRedisSerializer))
                // default TTL: 10 min
                .entryTtl(Duration.ofMinutes(10));

        // Different TTL configs
        Map<String, RedisCacheConfiguration> cacheConfigs = new HashMap<>();
        cacheConfigs.put("shortArticleListCache", defaultCacheConfig.entryTtl(Duration.ofMinutes(1)));
//        cacheConfigs.put("longLivedCache", defaultCacheConfig.entryTtl(Duration.ofHours(1)));

        return RedisCacheManager.builder(redisConnectionFactory)
                .cacheDefaults(defaultCacheConfig)
                .withInitialCacheConfigurations(cacheConfigs)
                .build(); // Builder design pattern
    }
}
