from os import environ


DURABLE: bool = True
AUTO_DELETE: bool = False
EXCLUSIVE: bool = False
AUTO_ACK: bool = False
MULTIPLE: bool = False
REQUEUE: bool = False
GLOBAL_PREFETCH: bool = True
PREFETCH_COUNT: int = int(environ.get("PREFETCH_COUNT", default=0))
PREFETCH_SIZE: int = 0
