import json
from datetime import datetime

import pytz
from fastapi import FastAPI
from fastapi.exceptions import RequestValidationError
from fastapi.middleware.cors import CORSMiddleware
from starlette import status
from starlette.concurrency import iterate_in_threadpool
from starlette.exceptions import HTTPException
from starlette.middleware.base import BaseHTTPMiddleware
from starlette.middleware.cors import CORSMiddleware
from starlette.requests import Request
from starlette.responses import JSONResponse


def add_fastapi_middleware(app: FastAPI):
    async def http_exception_handler(request, exc):
        return JSONResponse(
            status_code=exc.status_code,
            content={
                "status": exc.status_code,
                "message": exc.detail,
                "is_middleware_error": True,
            },
        )

    async def validation_exception_handler(request, exc):
        return JSONResponse(
            status_code=status.HTTP_200_OK,
            content={
                "status": 500,
                "message": "参数校验错误" + str(exc),
                "is_middleware_error": True,
            },
        )

    async def general_exception_handler(request, exc):
        return JSONResponse(  # pragma: no cover
            status_code=status.HTTP_200_OK,
            content={
                "status": 500,
                "message": "服务器错误" + str(exc),
                "timestamp": pytz.timezone("Asia/Shanghai")
                .localize(datetime.now())
                .strftime("%Y-%m-%d %H:%M:%S"),
            },
        )

    async def log_response(request: Request, call_next):
        response = await call_next(request)
        # return directly if not an api endpoint
        if not request.url.path.startswith("/api") or request.url.path.endswith(
            "access-token"
        ):
            return response
        response_body = [chunk async for chunk in response.body_iterator]
        response.body_iterator = iterate_in_threadpool(iter(response_body))
        response_body = json.loads(response_body[0].decode())
        if isinstance(response_body, dict) and response_body.get(
            "is_middleware_error", False
        ):
            return JSONResponse(
                status_code=status.HTTP_200_OK,
                content={
                    "status": response_body["status"],
                    "message": response_body["message"],
                    "timestamp": pytz.timezone("Asia/Shanghai")
                    .localize(datetime.now())
                    .strftime("%Y-%m-%d %H:%M:%S"),
                },
            )
        return JSONResponse(
            status_code=status.HTTP_200_OK,
            content={
                "status": response.status_code,
                "message": "success",
                "timestamp": pytz.timezone("Asia/Shanghai")
                .localize(datetime.now())
                .strftime("%Y-%m-%d %H:%M:%S"),
                "data": response_body,
            },
        )

    app.add_exception_handler(HTTPException, http_exception_handler)
    app.add_exception_handler(RequestValidationError, validation_exception_handler)
    app.add_middleware(BaseHTTPMiddleware, dispatch=log_response)
    app.add_middleware(
        CORSMiddleware,
        allow_origins=["*"],
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )
    app.add_exception_handler(Exception, general_exception_handler)
