-- Revert aldebaran:add_base_tables from pg

BEGIN;

DROP TABLE "rate";

DROP TABLE "selected_service";

DROP TABLE "job";

DROP TABLE "worker_service";

DROP TABLE "service";

DROP TABLE "schedule";

DROP TABLE "user_client";

DROP TABLE "user_worker";

DROP TABLE "date";


DROP TYPE "day_of_week_enum";

DROP TYPE "job_status_enum";

DROP TYPE "owner_type_enum";

DROP TYPE "gender_enum";

COMMIT;
