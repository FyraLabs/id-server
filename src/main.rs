use std::env;

use actix_web::{web::{self, Data}, get, App, HttpResponse, HttpServer, Responder};
use diesel::{PgConnection, r2d2::{ConnectionManager, Pool}, QueryDsl, RunQueryDsl};
use dotenv::dotenv;
use schema::users::dsl::users;

#[macro_use] extern crate diesel;

pub mod schema;
pub mod models;

type DbPool = Data<Pool<ConnectionManager<PgConnection>>>;

#[get("/")]
async fn index(pool: DbPool) -> impl Responder {
    HttpResponse::Ok().body("Hello!")
}

#[get("/users")]
async fn get_users(pool: DbPool) -> impl Responder {
    let other_users = users.load::<models::User>(&pool.get().expect("dasdsads"));
    HttpResponse::Ok().json(other_users.expect("dsadsads"))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv();
    
    let manager = ConnectionManager::<PgConnection>::new(env::var("DATABASE_URL").expect("Could not get database url"));
    let pool = diesel::r2d2::Pool::builder().build(manager).expect("Failed to create pool");

    HttpServer::new(move || App::new().app_data(pool.clone()).service(index).service(get_users))
        .bind(("127.0.0.1", 3000))?
        .run()
        .await
}
