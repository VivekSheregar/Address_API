BEGIN;
CREATE TABLE "mst_states"(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    CONSTRAINT "mst_state_name_ukey" UNIQUE ("name")
);

CREATE TABLE "mst_citys"(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "mst_state_id" INTEGER NOT  NULL,
    CONSTRAINT "mst_citys_mst_states_state_fkey" FOREIGN KEY ("mst_state_id") REFERENCES mst_states("id") ON DELETE CASCADE
);

COMMIT;