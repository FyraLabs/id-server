use diesel::Queryable;
use uuid::Uuid;
use chrono::{NaiveDateTime};
use serde::{Serialize};

#[derive(Queryable, Serialize)]
pub struct User {
    pub id: Uuid,
    pub name: String,
    pub email: String,
    pub password: String,
    pub created_at: NaiveDateTime,
}