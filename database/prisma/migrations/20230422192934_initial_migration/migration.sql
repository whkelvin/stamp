-- CreateTable
CREATE TABLE "posts" (
    "post_id" TEXT NOT NULL DEFAULT '',
    "link" TEXT NOT NULL DEFAULT '',
    "title" TEXT NOT NULL DEFAULT '',
    "description" TEXT NOT NULL DEFAULT '',
    "created_date" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "root_domain" TEXT NOT NULL DEFAULT '',

    CONSTRAINT "posts_pkey" PRIMARY KEY ("post_id")
);
