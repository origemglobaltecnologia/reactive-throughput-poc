package tech.origem.global.reactive;

import lombok.extern.slf4j.Slf4j;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Service;

@Service
@Slf4j
public class MessageConsumer {

    // Este método "esvazia" a fila assim que a mensagem chega
    @RabbitListener(queues = "high-throughput-queue", concurrency = "5-10")
    public void consume(MessageRequest message) {
        // Em um teste de carga, apenas logamos ou ignoramos para manter a velocidade
        // log.info("Mensagem processada da origem: {}", message.getOrigin());
    }
}
