package com.taranttini;

import com.taranttini.persistence.migration.MigrationStrategy;
import com.taranttini.ui.MainMenu;

import java.sql.SQLException;

import static com.taranttini.persistence.config.ConnectionConfig.getConnection;


public class Main {

    public static void main(String[] args) throws SQLException {

        new MigrationStrategy().executeMigration();

        new MainMenu().execute();
    }

}
