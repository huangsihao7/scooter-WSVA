import logging

from uvicorn.config import LOGGING_CONFIG


def setup_logger() -> dict:  # pragma: no cover
    """
    Cycles through uvicorn root loggers to
    remove handler, then setup logger config
    """

    # Remove all handlers from root logger
    # and proprogate to root logger.
    for name in logging.root.manager.loggerDict.keys():
        logging.getLogger(name).handlers = []
        logging.getLogger(name).propagate = True
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(name)s: %(message)s",
        datefmt="%Y-%m-%d-%H:%M:%S",
        handlers=[
            logging.StreamHandler(),
        ],
        force=True,
    )
    logging_config = LOGGING_CONFIG
    logging_config["loggers"] = {
        "uvicorn": {"level": "INFO"},
        "uvicorn.error": {"level": "INFO"},
        "uvicorn.access": {"level": "INFO"},
    }
    return logging_config
