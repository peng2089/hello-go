<?php
$cfg['blowfish_secret'] = '1{dd0`<Q),5XP_:R9UK%%8\"EEcyH#{o';

$cfg['Servers'] = [
    // localhost
    1 => [
        'auth_type' => 'config',
        'host' => 'mysql',
        'port' => 3306,
        'user' => 'root',
        'password' => '123456',
        'verbose' => 'localhost', // 详细名称
        // ------------------------------------------
        // phpMyAdmin configuration storage settings.
        // ------------------------------------------
        'controlhost' => 'mysql',
        'controlport' => '3306',
        'controluser' => 'pma',
        'controlpass' => 'r=PYdKnH)1%2bGIr',
        'pmadb' => 'phpmyadmin',
        'bookmarktable' => 'pma__bookmark',
        'relation' => 'pma__relation',
        'table_info' => 'pma__table_info',
        'table_coords' => 'pma__table_coords',
        'pdf_pages' => 'pma__pdf_pages',
        'column_info' => 'pma__column_info',
        'history' => 'pma__history',
        'table_uiprefs' => 'pma__table_uiprefs',
        'tracking' => 'pma__tracking',
        'userconfig' => 'pma__userconfig',
        'recent' => 'pma__recent',
        'favorite' => 'pma__favorite',
        'users' => 'pma__users',
        'usergroups' => 'pma__usergroups',
        'navigationhiding' => 'pma__navigationhiding',
        'savedsearches' => 'pma__savedsearches',
        'central_columns' => 'pma__central_columns',
        'designer_settings' => 'pma__designer_settings',
        'export_templates' => 'pma__export_templates',
    ],
];
