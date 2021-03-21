module github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}

go {{cookiecutter.docker_build_image_version}}

require (
	github.com/gofiber/adaptor/v2 v2.1.1
	github.com/gofiber/fiber/v2 v2.6.0
	github.com/prometheus/client_golang v1.7.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.7.1
	{% if cookiecutter.use_postgresql == "y" -%}
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jmoiron/sqlx v1.3.1
	github.com/pkg/errors v0.9.1
	{%- endif %}
	{% if cookiecutter.use_redis == "y" %}
	github.com/go-redis/redis/v8 v8.7.1
	{%- endif %}
)
