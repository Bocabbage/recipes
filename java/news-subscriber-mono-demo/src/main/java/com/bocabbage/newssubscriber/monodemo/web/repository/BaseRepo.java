package com.bocabbage.newssubscriber.monodemo.web.repository;

import com.bocabbage.newssubscriber.monodemo.web.entity.BaseEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.NoRepositoryBean;

import java.io.Serializable;

@NoRepositoryBean // Spring 无需生成对应的 Repo 实例
public interface BaseRepo<T extends BaseEntity, IdType extends Serializable> extends
                 JpaRepository<T, IdType>, JpaSpecificationExecutor<T>
{
    // JpaRepository 是 CrudRepository 的扩展，增加 Batch 操作等
    // JpaSpecificationExecutor 提供了基于 JPA Criteria API 的动态查询功能
}
