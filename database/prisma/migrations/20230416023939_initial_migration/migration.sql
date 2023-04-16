-- CreateTable
CREATE TABLE "post" (
    "post_id" SERIAL NOT NULL,
    "link" TEXT NOT NULL DEFAULT '',
    "title" TEXT NOT NULL DEFAULT '',
    "description" TEXT NOT NULL DEFAULT '',

    CONSTRAINT "post_pkey" PRIMARY KEY ("post_id")
);
