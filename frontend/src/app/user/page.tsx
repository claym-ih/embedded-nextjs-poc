"use client"

import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { useMany } from "@refinedev/core";
import {
    DateField,
    DeleteButton,
    EditButton,
    List,
    MarkdownField,
    ShowButton,
    useDataGrid,
} from "@refinedev/mui";
import React from "react";

export default function UserList() {
    const { dataGridProps } = useDataGrid({
        syncWithLocation: true,
    });

    const columns = React.useMemo<GridColDef[]>(
        () => [
            {
                field: "id",
                headerName: "ID",
                type: "number",
                minWidth: 50,
            },
            {
                field: "name",
                flex: 1,
                headerName: "Name",
                minWidth: 200,
            },
            {
                field: "email",
                flex: 1,
                headerName: "Email",
                minWidth: 250,
                renderCell: function render({ value }) {
                    if (!value) return "-";
                    return <MarkdownField value={value?.slice(0, 80) + "..." || ""} />;
                },
            },
            {
                field: "actions",
                headerName: "Actions",
                sortable: false,
                renderCell: function render({ row }) {
                    return (
                        <>
                            <EditButton hideText recordItemId={row.id} />
                            <ShowButton hideText recordItemId={row.id} />
                            <DeleteButton hideText recordItemId={row.id} />
                        </>
                    );
                },
                align: "center",
                headerAlign: "center",
                minWidth: 80,
            },
        ],
        []
    );

    return (
        <List>
            <DataGrid {...dataGridProps} columns={columns} autoHeight />
        </List>
    );
}
