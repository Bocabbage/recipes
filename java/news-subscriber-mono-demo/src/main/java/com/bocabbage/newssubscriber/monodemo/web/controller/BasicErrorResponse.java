package com.bocabbage.newssubscriber.monodemo.web.controller;

import lombok.Builder;
import lombok.Data;
import lombok.Singular;

import java.time.LocalDateTime;
import java.util.List;

@Data
@Builder // 提供 builder 方法
public class BasicErrorResponse {
    private String traceId;
    private String message;
    private LocalDateTime time;
    @Singular
    private List<Object> errors;
}
