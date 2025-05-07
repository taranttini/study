package com.taranttini.persistence.migration;

import com.taranttini.persistence.config.ConnectionConfig;
import lombok.AllArgsConstructor;
import org.flywaydb.core.Flyway;

@AllArgsConstructor
public class MigrationStrategy {

    public void executeMigration() {
        System.out.println("Starting migration...");
        var flyway = Flyway.configure()
                .dataSource(ConnectionConfig.getUrl(), ConnectionConfig.getUser(), ConnectionConfig.getPass())
                .load();
        flyway.migrate();

        System.out.println("Complete migration...");
    }
}
