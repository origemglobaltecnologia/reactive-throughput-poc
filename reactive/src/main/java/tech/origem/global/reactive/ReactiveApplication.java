package tech.origem.global.reactive;

import org.springframework.amqp.core.Queue;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
public class ReactiveApplication {

    public static void main(String[] args) {
        SpringApplication.run(ReactiveApplication.class, args);
    }

    @Bean
    public Queue highThroughputQueue() {
        // 'true' para durável, garantindo que apareça no LavinMQ
        return new Queue("high-throughput-queue", true);
    }
}
