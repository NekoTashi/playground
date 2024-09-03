FROM python:3.11.9-slim-bookworm

COPY --from=ghcr.io/astral-sh/uv:latest /uv /bin/uv

WORKDIR /app

ADD uv.lock /app/uv.lock
ADD pyproject.toml /app/pyproject.toml
RUN --mount=type=cache,target=/root/.cache/uv \
    uv sync --frozen --no-install-project

ADD . /app
RUN --mount=type=cache,target=/root/.cache/uv \
    uv sync --frozen

ENV PATH="/app/.venv/bin:$PATH"

ENV FLASK_APP=hello.py
CMD ["flask", "run", "-p", "8080"]
