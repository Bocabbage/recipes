package com.bocabbage.newssubscriber.monodemo.web.controller;

import jakarta.validation.ConstraintViolation;
import jakarta.validation.ConstraintViolationException;
import lombok.NonNull;
import lombok.extern.slf4j.Slf4j;
import org.springframework.context.MessageSourceResolvable;
import org.springframework.context.support.DefaultMessageSourceResolvable;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.http.converter.HttpMessageNotReadableException;
import org.springframework.validation.BindException;
import org.springframework.validation.BindingResult;
import org.springframework.web.HttpRequestMethodNotSupportedException;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.method.annotation.HandlerMethodValidationException;
import org.springframework.web.method.annotation.MethodArgumentTypeMismatchException;
import org.springframework.web.servlet.resource.NoResourceFoundException;

import java.time.LocalDateTime;
import java.util.HashMap;
import java.util.UUID;


@Slf4j
@RestControllerAdvice // 全局处理控制器异常，绑定数据，提供全局数据等
public class GlobalExceptionHandler {

    // 通用的 Client 问题
    @ExceptionHandler({
            BindException.class,
            NoResourceFoundException.class,
            MethodArgumentTypeMismatchException.class,
            ConstraintViolationException.class,
            HandlerMethodValidationException.class,
            HttpMessageNotReadableException.class,
    })
    public ResponseEntity<BasicErrorResponse> clientExceptionHandler(@NonNull Exception e) {
        var uuid = UUID.randomUUID().toString();
        var basicErrorRespBuilder = BasicErrorResponse.builder()
                .traceId(uuid)
                .time(LocalDateTime.now());

        if(e instanceof HttpRequestMethodNotSupportedException) {
            log.warn("[HttpRequestMethodNotSupportedException][traceId: " + uuid +"]: ", e);
            basicErrorRespBuilder.message("Method not valid.");
            return ResponseEntity.status(HttpStatus.METHOD_NOT_ALLOWED).body(basicErrorRespBuilder.build());
        }
        else if(e instanceof BindException) {
            log.warn("[BindException][traceId: " + uuid + "]: ", e);
            BindingResult bindingResult = ((BindException) e).getBindingResult();
            bindingResult.getAllErrors().stream()
                    .map(DefaultMessageSourceResolvable::getDefaultMessage)
                    .forEach(basicErrorRespBuilder::error);

            basicErrorRespBuilder.message("Bind argument invalid");
        }
        else if(e instanceof HttpMessageNotReadableException) {
            basicErrorRespBuilder.message("HTTP message not readable: format error");
        }
        else if(e instanceof NoResourceFoundException) {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body(basicErrorRespBuilder.message("Resource not found").build());
        }
//        else if(e instanceof ConstraintViolationException) {
//            var violationResult = ((ConstraintViolationException) e).getConstraintViolations();
//            violationResult.stream()
//                    .map(ConstraintViolation::getMessage)
//                    .forEach(basicErrorRespBuilder::error);
//            basicErrorRespBuilder.message("Path argument invalid");
//        }
        else if(e instanceof HandlerMethodValidationException) {
            var exResult = ((HandlerMethodValidationException) e).getAllErrors();
            exResult.stream()
                    .map(MessageSourceResolvable::getDefaultMessage)
                    .forEach(basicErrorRespBuilder::error);
            // TODO: maybe enhanced with field name.
            basicErrorRespBuilder.message("Path argument invalid");
        }
        else {
            log.warn("[MethodArgumentNotValidException][traceId: " + uuid + "]: ", e);
            basicErrorRespBuilder
                    .message("Method argument invalid")
                    .error(((MethodArgumentTypeMismatchException) e).getName());
        }
        return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(basicErrorRespBuilder.build());

    }

    @ExceptionHandler(Exception.class)
    @ResponseStatus(HttpStatus.INTERNAL_SERVER_ERROR)
    public BasicErrorResponse unexpectedExceptionHandler(@NonNull Exception e) {
        var uuid = UUID.randomUUID().toString();
        var basicErrorRespBuilder = BasicErrorResponse.builder();
        log.error("[UnexpectedError][traceId: " + uuid + "]: ", e);
        return basicErrorRespBuilder
                .traceId(uuid)
                .message("Internal Error happened.")
                .time(LocalDateTime.now())
                .build();
    }
}
