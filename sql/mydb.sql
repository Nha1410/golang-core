-- 必ず勤務する曜日
-- * BackupToTempTable
DROP TABLE if exists m__pets CASCADE;

-- * RestoreFromTempTable
CREATE TABLE m__pets (
  id uuid PRIMARY KEY,
    name text NOT NULL,
    age int NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
    created_by uuid NOT NULL,
    updated_by uuid NOT NULL
    lock_version int8 NOT NULL DEFAULT 0
);


