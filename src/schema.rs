table! {
    users (id) {
        id -> Uuid,
        name -> Text,
        email -> Text,
        password -> Text,
        created_at -> Timestamp,
    }
}
