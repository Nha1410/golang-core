-- * BackupToTempTable
DROP TABLE IF EXISTS m__pets CASCADE;

CREATE TABLE m__pets (
	id uuid DEFAULT gen_random_uuid() NOT NULL,
	"name" text NOT NULL,
	age int4 NOT NULL,
	status text NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp DEFAULT now() NOT NULL,
	created_by uuid NOT NULL,
	updated_by uuid NOT NULL,
	lock_version int8 DEFAULT 0 NOT NULL,
	CONSTRAINT m__pets_pkey PRIMARY KEY (id)
);
