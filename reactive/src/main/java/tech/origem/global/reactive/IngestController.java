package tech.origem.global.reactive;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Mono;
import reactor.core.scheduler.Schedulers;

@RestController
@RequiredArgsConstructor
@Slf4j
public class IngestController {

    private final RabbitTemplate rabbitTemplate;
    private static final String QUEUE_NAME = "high-throughput-queue";

    @PostMapping("/ingest")
    public Mono<ResponseEntity<Void>> ingest(@RequestBody MessageRequest request) {
        return Mono.fromRunnable(() -> 
            rabbitTemplate.convertAndSend(QUEUE_NAME, request)
        )
        .subscribeOn(Schedulers.boundedElastic()) // Não bloqueia o Event Loop do Netty
        .thenReturn(new ResponseEntity<Void>(HttpStatus.ACCEPTED));
    }
}
