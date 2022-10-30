package com.LMS.mapper;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.autoconfigure.SpringBootApplication;

/**
 * <p>
 * </p>
 *
 * @author xxz
 * @date 2022/10/30 14:31
 */

@MapperScan("com.lms.mapper")
@SpringBootApplication
public class DemoApplication {
    public static void main(String[] args) {
        SpringBootApplication.run(DemoApplication.class,args);
    }
}
