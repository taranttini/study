package com.taranttini.persistence.config;

import lombok.NoArgsConstructor;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;

import static lombok.AccessLevel.PRIVATE;

@NoArgsConstructor(access = PRIVATE)
public final class ConnectionConfig {

    // changes to get by env
    public static String getUser() { return "java-board"; }
    public static String getPass() { return "java-board"; }
    public static String getUrl() { return "jdbc:mysql://localhost/java-board"; }

    public static Connection getConnection() throws SQLException {
        var connection = DriverManager.getConnection(getUrl(), getUser(), getPass());
        connection.setAutoCommit(false);
        return connection;
    }

}
