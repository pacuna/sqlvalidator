// Code generated from HiveParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // HiveParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HiveParserListener is a complete listener for a parse tree produced by HiveParser.
type HiveParserListener interface {
	antlr.ParseTreeListener

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterExplainStatement is called when entering the explainStatement production.
	EnterExplainStatement(c *ExplainStatementContext)

	// EnterExplainOption is called when entering the explainOption production.
	EnterExplainOption(c *ExplainOptionContext)

	// EnterVectorizationOnly is called when entering the vectorizationOnly production.
	EnterVectorizationOnly(c *VectorizationOnlyContext)

	// EnterVectorizatonDetail is called when entering the vectorizatonDetail production.
	EnterVectorizatonDetail(c *VectorizatonDetailContext)

	// EnterExecStatement is called when entering the execStatement production.
	EnterExecStatement(c *ExecStatementContext)

	// EnterLoadStatement is called when entering the loadStatement production.
	EnterLoadStatement(c *LoadStatementContext)

	// EnterReplicationClause is called when entering the replicationClause production.
	EnterReplicationClause(c *ReplicationClauseContext)

	// EnterExportStatement is called when entering the exportStatement production.
	EnterExportStatement(c *ExportStatementContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterReplDumpStatement is called when entering the replDumpStatement production.
	EnterReplDumpStatement(c *ReplDumpStatementContext)

	// EnterReplLoadStatement is called when entering the replLoadStatement production.
	EnterReplLoadStatement(c *ReplLoadStatementContext)

	// EnterReplConfigs is called when entering the replConfigs production.
	EnterReplConfigs(c *ReplConfigsContext)

	// EnterReplConfigsList is called when entering the replConfigsList production.
	EnterReplConfigsList(c *ReplConfigsListContext)

	// EnterReplStatusStatement is called when entering the replStatusStatement production.
	EnterReplStatusStatement(c *ReplStatusStatementContext)

	// EnterDdlStatement is called when entering the ddlStatement production.
	EnterDdlStatement(c *DdlStatementContext)

	// EnterIfExists is called when entering the ifExists production.
	EnterIfExists(c *IfExistsContext)

	// EnterRestrictOrCascade is called when entering the restrictOrCascade production.
	EnterRestrictOrCascade(c *RestrictOrCascadeContext)

	// EnterIfNotExists is called when entering the ifNotExists production.
	EnterIfNotExists(c *IfNotExistsContext)

	// EnterRewriteEnabled is called when entering the rewriteEnabled production.
	EnterRewriteEnabled(c *RewriteEnabledContext)

	// EnterRewriteDisabled is called when entering the rewriteDisabled production.
	EnterRewriteDisabled(c *RewriteDisabledContext)

	// EnterStoredAsDirs is called when entering the storedAsDirs production.
	EnterStoredAsDirs(c *StoredAsDirsContext)

	// EnterOrReplace is called when entering the orReplace production.
	EnterOrReplace(c *OrReplaceContext)

	// EnterCreateDatabaseStatement is called when entering the createDatabaseStatement production.
	EnterCreateDatabaseStatement(c *CreateDatabaseStatementContext)

	// EnterDbLocation is called when entering the dbLocation production.
	EnterDbLocation(c *DbLocationContext)

	// EnterDbProperties is called when entering the dbProperties production.
	EnterDbProperties(c *DbPropertiesContext)

	// EnterDbPropertiesList is called when entering the dbPropertiesList production.
	EnterDbPropertiesList(c *DbPropertiesListContext)

	// EnterSwitchDatabaseStatement is called when entering the switchDatabaseStatement production.
	EnterSwitchDatabaseStatement(c *SwitchDatabaseStatementContext)

	// EnterDropDatabaseStatement is called when entering the dropDatabaseStatement production.
	EnterDropDatabaseStatement(c *DropDatabaseStatementContext)

	// EnterDatabaseComment is called when entering the databaseComment production.
	EnterDatabaseComment(c *DatabaseCommentContext)

	// EnterCreateTableStatement is called when entering the createTableStatement production.
	EnterCreateTableStatement(c *CreateTableStatementContext)

	// EnterTruncateTableStatement is called when entering the truncateTableStatement production.
	EnterTruncateTableStatement(c *TruncateTableStatementContext)

	// EnterDropTableStatement is called when entering the dropTableStatement production.
	EnterDropTableStatement(c *DropTableStatementContext)

	// EnterAlterStatement is called when entering the alterStatement production.
	EnterAlterStatement(c *AlterStatementContext)

	// EnterAlterTableStatementSuffix is called when entering the alterTableStatementSuffix production.
	EnterAlterTableStatementSuffix(c *AlterTableStatementSuffixContext)

	// EnterAlterTblPartitionStatementSuffix is called when entering the alterTblPartitionStatementSuffix production.
	EnterAlterTblPartitionStatementSuffix(c *AlterTblPartitionStatementSuffixContext)

	// EnterAlterStatementPartitionKeyType is called when entering the alterStatementPartitionKeyType production.
	EnterAlterStatementPartitionKeyType(c *AlterStatementPartitionKeyTypeContext)

	// EnterAlterViewStatementSuffix is called when entering the alterViewStatementSuffix production.
	EnterAlterViewStatementSuffix(c *AlterViewStatementSuffixContext)

	// EnterAlterMaterializedViewStatementSuffix is called when entering the alterMaterializedViewStatementSuffix production.
	EnterAlterMaterializedViewStatementSuffix(c *AlterMaterializedViewStatementSuffixContext)

	// EnterAlterDatabaseStatementSuffix is called when entering the alterDatabaseStatementSuffix production.
	EnterAlterDatabaseStatementSuffix(c *AlterDatabaseStatementSuffixContext)

	// EnterAlterDatabaseSuffixProperties is called when entering the alterDatabaseSuffixProperties production.
	EnterAlterDatabaseSuffixProperties(c *AlterDatabaseSuffixPropertiesContext)

	// EnterAlterDatabaseSuffixSetOwner is called when entering the alterDatabaseSuffixSetOwner production.
	EnterAlterDatabaseSuffixSetOwner(c *AlterDatabaseSuffixSetOwnerContext)

	// EnterAlterDatabaseSuffixSetLocation is called when entering the alterDatabaseSuffixSetLocation production.
	EnterAlterDatabaseSuffixSetLocation(c *AlterDatabaseSuffixSetLocationContext)

	// EnterAlterStatementSuffixRename is called when entering the alterStatementSuffixRename production.
	EnterAlterStatementSuffixRename(c *AlterStatementSuffixRenameContext)

	// EnterAlterStatementSuffixAddCol is called when entering the alterStatementSuffixAddCol production.
	EnterAlterStatementSuffixAddCol(c *AlterStatementSuffixAddColContext)

	// EnterAlterStatementSuffixAddConstraint is called when entering the alterStatementSuffixAddConstraint production.
	EnterAlterStatementSuffixAddConstraint(c *AlterStatementSuffixAddConstraintContext)

	// EnterAlterStatementSuffixUpdateColumns is called when entering the alterStatementSuffixUpdateColumns production.
	EnterAlterStatementSuffixUpdateColumns(c *AlterStatementSuffixUpdateColumnsContext)

	// EnterAlterStatementSuffixDropConstraint is called when entering the alterStatementSuffixDropConstraint production.
	EnterAlterStatementSuffixDropConstraint(c *AlterStatementSuffixDropConstraintContext)

	// EnterAlterStatementSuffixRenameCol is called when entering the alterStatementSuffixRenameCol production.
	EnterAlterStatementSuffixRenameCol(c *AlterStatementSuffixRenameColContext)

	// EnterAlterStatementSuffixUpdateStatsCol is called when entering the alterStatementSuffixUpdateStatsCol production.
	EnterAlterStatementSuffixUpdateStatsCol(c *AlterStatementSuffixUpdateStatsColContext)

	// EnterAlterStatementSuffixUpdateStats is called when entering the alterStatementSuffixUpdateStats production.
	EnterAlterStatementSuffixUpdateStats(c *AlterStatementSuffixUpdateStatsContext)

	// EnterAlterStatementChangeColPosition is called when entering the alterStatementChangeColPosition production.
	EnterAlterStatementChangeColPosition(c *AlterStatementChangeColPositionContext)

	// EnterAlterStatementSuffixAddPartitions is called when entering the alterStatementSuffixAddPartitions production.
	EnterAlterStatementSuffixAddPartitions(c *AlterStatementSuffixAddPartitionsContext)

	// EnterAlterStatementSuffixAddPartitionsElement is called when entering the alterStatementSuffixAddPartitionsElement production.
	EnterAlterStatementSuffixAddPartitionsElement(c *AlterStatementSuffixAddPartitionsElementContext)

	// EnterAlterStatementSuffixTouch is called when entering the alterStatementSuffixTouch production.
	EnterAlterStatementSuffixTouch(c *AlterStatementSuffixTouchContext)

	// EnterAlterStatementSuffixArchive is called when entering the alterStatementSuffixArchive production.
	EnterAlterStatementSuffixArchive(c *AlterStatementSuffixArchiveContext)

	// EnterAlterStatementSuffixUnArchive is called when entering the alterStatementSuffixUnArchive production.
	EnterAlterStatementSuffixUnArchive(c *AlterStatementSuffixUnArchiveContext)

	// EnterPartitionLocation is called when entering the partitionLocation production.
	EnterPartitionLocation(c *PartitionLocationContext)

	// EnterAlterStatementSuffixDropPartitions is called when entering the alterStatementSuffixDropPartitions production.
	EnterAlterStatementSuffixDropPartitions(c *AlterStatementSuffixDropPartitionsContext)

	// EnterAlterStatementSuffixProperties is called when entering the alterStatementSuffixProperties production.
	EnterAlterStatementSuffixProperties(c *AlterStatementSuffixPropertiesContext)

	// EnterAlterViewSuffixProperties is called when entering the alterViewSuffixProperties production.
	EnterAlterViewSuffixProperties(c *AlterViewSuffixPropertiesContext)

	// EnterAlterMaterializedViewSuffixRewrite is called when entering the alterMaterializedViewSuffixRewrite production.
	EnterAlterMaterializedViewSuffixRewrite(c *AlterMaterializedViewSuffixRewriteContext)

	// EnterAlterMaterializedViewSuffixRebuild is called when entering the alterMaterializedViewSuffixRebuild production.
	EnterAlterMaterializedViewSuffixRebuild(c *AlterMaterializedViewSuffixRebuildContext)

	// EnterAlterStatementSuffixSerdeProperties is called when entering the alterStatementSuffixSerdeProperties production.
	EnterAlterStatementSuffixSerdeProperties(c *AlterStatementSuffixSerdePropertiesContext)

	// EnterTablePartitionPrefix is called when entering the tablePartitionPrefix production.
	EnterTablePartitionPrefix(c *TablePartitionPrefixContext)

	// EnterAlterStatementSuffixFileFormat is called when entering the alterStatementSuffixFileFormat production.
	EnterAlterStatementSuffixFileFormat(c *AlterStatementSuffixFileFormatContext)

	// EnterAlterStatementSuffixClusterbySortby is called when entering the alterStatementSuffixClusterbySortby production.
	EnterAlterStatementSuffixClusterbySortby(c *AlterStatementSuffixClusterbySortbyContext)

	// EnterAlterTblPartitionStatementSuffixSkewedLocation is called when entering the alterTblPartitionStatementSuffixSkewedLocation production.
	EnterAlterTblPartitionStatementSuffixSkewedLocation(c *AlterTblPartitionStatementSuffixSkewedLocationContext)

	// EnterSkewedLocations is called when entering the skewedLocations production.
	EnterSkewedLocations(c *SkewedLocationsContext)

	// EnterSkewedLocationsList is called when entering the skewedLocationsList production.
	EnterSkewedLocationsList(c *SkewedLocationsListContext)

	// EnterSkewedLocationMap is called when entering the skewedLocationMap production.
	EnterSkewedLocationMap(c *SkewedLocationMapContext)

	// EnterAlterStatementSuffixLocation is called when entering the alterStatementSuffixLocation production.
	EnterAlterStatementSuffixLocation(c *AlterStatementSuffixLocationContext)

	// EnterAlterStatementSuffixSkewedby is called when entering the alterStatementSuffixSkewedby production.
	EnterAlterStatementSuffixSkewedby(c *AlterStatementSuffixSkewedbyContext)

	// EnterAlterStatementSuffixExchangePartition is called when entering the alterStatementSuffixExchangePartition production.
	EnterAlterStatementSuffixExchangePartition(c *AlterStatementSuffixExchangePartitionContext)

	// EnterAlterStatementSuffixRenamePart is called when entering the alterStatementSuffixRenamePart production.
	EnterAlterStatementSuffixRenamePart(c *AlterStatementSuffixRenamePartContext)

	// EnterAlterStatementSuffixStatsPart is called when entering the alterStatementSuffixStatsPart production.
	EnterAlterStatementSuffixStatsPart(c *AlterStatementSuffixStatsPartContext)

	// EnterAlterStatementSuffixMergeFiles is called when entering the alterStatementSuffixMergeFiles production.
	EnterAlterStatementSuffixMergeFiles(c *AlterStatementSuffixMergeFilesContext)

	// EnterAlterStatementSuffixBucketNum is called when entering the alterStatementSuffixBucketNum production.
	EnterAlterStatementSuffixBucketNum(c *AlterStatementSuffixBucketNumContext)

	// EnterBlocking is called when entering the blocking production.
	EnterBlocking(c *BlockingContext)

	// EnterAlterStatementSuffixCompact is called when entering the alterStatementSuffixCompact production.
	EnterAlterStatementSuffixCompact(c *AlterStatementSuffixCompactContext)

	// EnterAlterStatementSuffixSetOwner is called when entering the alterStatementSuffixSetOwner production.
	EnterAlterStatementSuffixSetOwner(c *AlterStatementSuffixSetOwnerContext)

	// EnterFileFormat is called when entering the fileFormat production.
	EnterFileFormat(c *FileFormatContext)

	// EnterInputFileFormat is called when entering the inputFileFormat production.
	EnterInputFileFormat(c *InputFileFormatContext)

	// EnterTabTypeExpr is called when entering the tabTypeExpr production.
	EnterTabTypeExpr(c *TabTypeExprContext)

	// EnterPartTypeExpr is called when entering the partTypeExpr production.
	EnterPartTypeExpr(c *PartTypeExprContext)

	// EnterTabPartColTypeExpr is called when entering the tabPartColTypeExpr production.
	EnterTabPartColTypeExpr(c *TabPartColTypeExprContext)

	// EnterDescStatement is called when entering the descStatement production.
	EnterDescStatement(c *DescStatementContext)

	// EnterAnalyzeStatement is called when entering the analyzeStatement production.
	EnterAnalyzeStatement(c *AnalyzeStatementContext)

	// EnterShowStatement is called when entering the showStatement production.
	EnterShowStatement(c *ShowStatementContext)

	// EnterLockStatement is called when entering the lockStatement production.
	EnterLockStatement(c *LockStatementContext)

	// EnterLockDatabase is called when entering the lockDatabase production.
	EnterLockDatabase(c *LockDatabaseContext)

	// EnterLockMode is called when entering the lockMode production.
	EnterLockMode(c *LockModeContext)

	// EnterUnlockStatement is called when entering the unlockStatement production.
	EnterUnlockStatement(c *UnlockStatementContext)

	// EnterUnlockDatabase is called when entering the unlockDatabase production.
	EnterUnlockDatabase(c *UnlockDatabaseContext)

	// EnterCreateRoleStatement is called when entering the createRoleStatement production.
	EnterCreateRoleStatement(c *CreateRoleStatementContext)

	// EnterDropRoleStatement is called when entering the dropRoleStatement production.
	EnterDropRoleStatement(c *DropRoleStatementContext)

	// EnterGrantPrivileges is called when entering the grantPrivileges production.
	EnterGrantPrivileges(c *GrantPrivilegesContext)

	// EnterRevokePrivileges is called when entering the revokePrivileges production.
	EnterRevokePrivileges(c *RevokePrivilegesContext)

	// EnterGrantRole is called when entering the grantRole production.
	EnterGrantRole(c *GrantRoleContext)

	// EnterRevokeRole is called when entering the revokeRole production.
	EnterRevokeRole(c *RevokeRoleContext)

	// EnterShowRoleGrants is called when entering the showRoleGrants production.
	EnterShowRoleGrants(c *ShowRoleGrantsContext)

	// EnterShowRoles is called when entering the showRoles production.
	EnterShowRoles(c *ShowRolesContext)

	// EnterShowCurrentRole is called when entering the showCurrentRole production.
	EnterShowCurrentRole(c *ShowCurrentRoleContext)

	// EnterSetRole is called when entering the setRole production.
	EnterSetRole(c *SetRoleContext)

	// EnterShowGrants is called when entering the showGrants production.
	EnterShowGrants(c *ShowGrantsContext)

	// EnterShowRolePrincipals is called when entering the showRolePrincipals production.
	EnterShowRolePrincipals(c *ShowRolePrincipalsContext)

	// EnterPrivilegeIncludeColObject is called when entering the privilegeIncludeColObject production.
	EnterPrivilegeIncludeColObject(c *PrivilegeIncludeColObjectContext)

	// EnterPrivilegeObject is called when entering the privilegeObject production.
	EnterPrivilegeObject(c *PrivilegeObjectContext)

	// EnterPrivObject is called when entering the privObject production.
	EnterPrivObject(c *PrivObjectContext)

	// EnterPrivObjectCols is called when entering the privObjectCols production.
	EnterPrivObjectCols(c *PrivObjectColsContext)

	// EnterPrivilegeList is called when entering the privilegeList production.
	EnterPrivilegeList(c *PrivilegeListContext)

	// EnterPrivlegeDef is called when entering the privlegeDef production.
	EnterPrivlegeDef(c *PrivlegeDefContext)

	// EnterPrivilegeType is called when entering the privilegeType production.
	EnterPrivilegeType(c *PrivilegeTypeContext)

	// EnterPrincipalSpecification is called when entering the principalSpecification production.
	EnterPrincipalSpecification(c *PrincipalSpecificationContext)

	// EnterPrincipalName is called when entering the principalName production.
	EnterPrincipalName(c *PrincipalNameContext)

	// EnterWithGrantOption is called when entering the withGrantOption production.
	EnterWithGrantOption(c *WithGrantOptionContext)

	// EnterGrantOptionFor is called when entering the grantOptionFor production.
	EnterGrantOptionFor(c *GrantOptionForContext)

	// EnterAdminOptionFor is called when entering the adminOptionFor production.
	EnterAdminOptionFor(c *AdminOptionForContext)

	// EnterWithAdminOption is called when entering the withAdminOption production.
	EnterWithAdminOption(c *WithAdminOptionContext)

	// EnterMetastoreCheck is called when entering the metastoreCheck production.
	EnterMetastoreCheck(c *MetastoreCheckContext)

	// EnterResourceList is called when entering the resourceList production.
	EnterResourceList(c *ResourceListContext)

	// EnterResource is called when entering the resource production.
	EnterResource(c *ResourceContext)

	// EnterResourceType is called when entering the resourceType production.
	EnterResourceType(c *ResourceTypeContext)

	// EnterCreateFunctionStatement is called when entering the createFunctionStatement production.
	EnterCreateFunctionStatement(c *CreateFunctionStatementContext)

	// EnterDropFunctionStatement is called when entering the dropFunctionStatement production.
	EnterDropFunctionStatement(c *DropFunctionStatementContext)

	// EnterReloadFunctionStatement is called when entering the reloadFunctionStatement production.
	EnterReloadFunctionStatement(c *ReloadFunctionStatementContext)

	// EnterCreateMacroStatement is called when entering the createMacroStatement production.
	EnterCreateMacroStatement(c *CreateMacroStatementContext)

	// EnterDropMacroStatement is called when entering the dropMacroStatement production.
	EnterDropMacroStatement(c *DropMacroStatementContext)

	// EnterCreateViewStatement is called when entering the createViewStatement production.
	EnterCreateViewStatement(c *CreateViewStatementContext)

	// EnterCreateMaterializedViewStatement is called when entering the createMaterializedViewStatement production.
	EnterCreateMaterializedViewStatement(c *CreateMaterializedViewStatementContext)

	// EnterViewPartition is called when entering the viewPartition production.
	EnterViewPartition(c *ViewPartitionContext)

	// EnterDropViewStatement is called when entering the dropViewStatement production.
	EnterDropViewStatement(c *DropViewStatementContext)

	// EnterDropMaterializedViewStatement is called when entering the dropMaterializedViewStatement production.
	EnterDropMaterializedViewStatement(c *DropMaterializedViewStatementContext)

	// EnterShowFunctionIdentifier is called when entering the showFunctionIdentifier production.
	EnterShowFunctionIdentifier(c *ShowFunctionIdentifierContext)

	// EnterShowStmtIdentifier is called when entering the showStmtIdentifier production.
	EnterShowStmtIdentifier(c *ShowStmtIdentifierContext)

	// EnterTableComment is called when entering the tableComment production.
	EnterTableComment(c *TableCommentContext)

	// EnterTablePartition is called when entering the tablePartition production.
	EnterTablePartition(c *TablePartitionContext)

	// EnterTableBuckets is called when entering the tableBuckets production.
	EnterTableBuckets(c *TableBucketsContext)

	// EnterTableSkewed is called when entering the tableSkewed production.
	EnterTableSkewed(c *TableSkewedContext)

	// EnterRowFormat is called when entering the rowFormat production.
	EnterRowFormat(c *RowFormatContext)

	// EnterRecordReader is called when entering the recordReader production.
	EnterRecordReader(c *RecordReaderContext)

	// EnterRecordWriter is called when entering the recordWriter production.
	EnterRecordWriter(c *RecordWriterContext)

	// EnterRowFormatSerde is called when entering the rowFormatSerde production.
	EnterRowFormatSerde(c *RowFormatSerdeContext)

	// EnterRowFormatDelimited is called when entering the rowFormatDelimited production.
	EnterRowFormatDelimited(c *RowFormatDelimitedContext)

	// EnterTableRowFormat is called when entering the tableRowFormat production.
	EnterTableRowFormat(c *TableRowFormatContext)

	// EnterTablePropertiesPrefixed is called when entering the tablePropertiesPrefixed production.
	EnterTablePropertiesPrefixed(c *TablePropertiesPrefixedContext)

	// EnterTableProperties is called when entering the tableProperties production.
	EnterTableProperties(c *TablePropertiesContext)

	// EnterTablePropertiesList is called when entering the tablePropertiesList production.
	EnterTablePropertiesList(c *TablePropertiesListContext)

	// EnterKeyValueProperty is called when entering the keyValueProperty production.
	EnterKeyValueProperty(c *KeyValuePropertyContext)

	// EnterKeyProperty is called when entering the keyProperty production.
	EnterKeyProperty(c *KeyPropertyContext)

	// EnterTableRowFormatFieldIdentifier is called when entering the tableRowFormatFieldIdentifier production.
	EnterTableRowFormatFieldIdentifier(c *TableRowFormatFieldIdentifierContext)

	// EnterTableRowFormatCollItemsIdentifier is called when entering the tableRowFormatCollItemsIdentifier production.
	EnterTableRowFormatCollItemsIdentifier(c *TableRowFormatCollItemsIdentifierContext)

	// EnterTableRowFormatMapKeysIdentifier is called when entering the tableRowFormatMapKeysIdentifier production.
	EnterTableRowFormatMapKeysIdentifier(c *TableRowFormatMapKeysIdentifierContext)

	// EnterTableRowFormatLinesIdentifier is called when entering the tableRowFormatLinesIdentifier production.
	EnterTableRowFormatLinesIdentifier(c *TableRowFormatLinesIdentifierContext)

	// EnterTableRowNullFormat is called when entering the tableRowNullFormat production.
	EnterTableRowNullFormat(c *TableRowNullFormatContext)

	// EnterTableFileFormat is called when entering the tableFileFormat production.
	EnterTableFileFormat(c *TableFileFormatContext)

	// EnterTableLocation is called when entering the tableLocation production.
	EnterTableLocation(c *TableLocationContext)

	// EnterColumnNameTypeList is called when entering the columnNameTypeList production.
	EnterColumnNameTypeList(c *ColumnNameTypeListContext)

	// EnterColumnNameTypeOrConstraintList is called when entering the columnNameTypeOrConstraintList production.
	EnterColumnNameTypeOrConstraintList(c *ColumnNameTypeOrConstraintListContext)

	// EnterColumnNameColonTypeList is called when entering the columnNameColonTypeList production.
	EnterColumnNameColonTypeList(c *ColumnNameColonTypeListContext)

	// EnterColumnNameList is called when entering the columnNameList production.
	EnterColumnNameList(c *ColumnNameListContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterExtColumnName is called when entering the extColumnName production.
	EnterExtColumnName(c *ExtColumnNameContext)

	// EnterColumnNameOrderList is called when entering the columnNameOrderList production.
	EnterColumnNameOrderList(c *ColumnNameOrderListContext)

	// EnterColumnParenthesesList is called when entering the columnParenthesesList production.
	EnterColumnParenthesesList(c *ColumnParenthesesListContext)

	// EnterEnableValidateSpecification is called when entering the enableValidateSpecification production.
	EnterEnableValidateSpecification(c *EnableValidateSpecificationContext)

	// EnterEnableSpecification is called when entering the enableSpecification production.
	EnterEnableSpecification(c *EnableSpecificationContext)

	// EnterValidateSpecification is called when entering the validateSpecification production.
	EnterValidateSpecification(c *ValidateSpecificationContext)

	// EnterEnforcedSpecification is called when entering the enforcedSpecification production.
	EnterEnforcedSpecification(c *EnforcedSpecificationContext)

	// EnterRelySpecification is called when entering the relySpecification production.
	EnterRelySpecification(c *RelySpecificationContext)

	// EnterCreateConstraint is called when entering the createConstraint production.
	EnterCreateConstraint(c *CreateConstraintContext)

	// EnterAlterConstraintWithName is called when entering the alterConstraintWithName production.
	EnterAlterConstraintWithName(c *AlterConstraintWithNameContext)

	// EnterTableLevelConstraint is called when entering the tableLevelConstraint production.
	EnterTableLevelConstraint(c *TableLevelConstraintContext)

	// EnterPkUkConstraint is called when entering the pkUkConstraint production.
	EnterPkUkConstraint(c *PkUkConstraintContext)

	// EnterCheckConstraint is called when entering the checkConstraint production.
	EnterCheckConstraint(c *CheckConstraintContext)

	// EnterCreateForeignKey is called when entering the createForeignKey production.
	EnterCreateForeignKey(c *CreateForeignKeyContext)

	// EnterAlterForeignKeyWithName is called when entering the alterForeignKeyWithName production.
	EnterAlterForeignKeyWithName(c *AlterForeignKeyWithNameContext)

	// EnterSkewedValueElement is called when entering the skewedValueElement production.
	EnterSkewedValueElement(c *SkewedValueElementContext)

	// EnterSkewedColumnValuePairList is called when entering the skewedColumnValuePairList production.
	EnterSkewedColumnValuePairList(c *SkewedColumnValuePairListContext)

	// EnterSkewedColumnValuePair is called when entering the skewedColumnValuePair production.
	EnterSkewedColumnValuePair(c *SkewedColumnValuePairContext)

	// EnterSkewedColumnValues is called when entering the skewedColumnValues production.
	EnterSkewedColumnValues(c *SkewedColumnValuesContext)

	// EnterSkewedColumnValue is called when entering the skewedColumnValue production.
	EnterSkewedColumnValue(c *SkewedColumnValueContext)

	// EnterSkewedValueLocationElement is called when entering the skewedValueLocationElement production.
	EnterSkewedValueLocationElement(c *SkewedValueLocationElementContext)

	// EnterOrderSpecification is called when entering the orderSpecification production.
	EnterOrderSpecification(c *OrderSpecificationContext)

	// EnterNullOrdering is called when entering the nullOrdering production.
	EnterNullOrdering(c *NullOrderingContext)

	// EnterColumnNameOrder is called when entering the columnNameOrder production.
	EnterColumnNameOrder(c *ColumnNameOrderContext)

	// EnterColumnNameCommentList is called when entering the columnNameCommentList production.
	EnterColumnNameCommentList(c *ColumnNameCommentListContext)

	// EnterColumnNameComment is called when entering the columnNameComment production.
	EnterColumnNameComment(c *ColumnNameCommentContext)

	// EnterColumnRefOrder is called when entering the columnRefOrder production.
	EnterColumnRefOrder(c *ColumnRefOrderContext)

	// EnterColumnNameType is called when entering the columnNameType production.
	EnterColumnNameType(c *ColumnNameTypeContext)

	// EnterColumnNameTypeOrConstraint is called when entering the columnNameTypeOrConstraint production.
	EnterColumnNameTypeOrConstraint(c *ColumnNameTypeOrConstraintContext)

	// EnterTableConstraint is called when entering the tableConstraint production.
	EnterTableConstraint(c *TableConstraintContext)

	// EnterColumnNameTypeConstraint is called when entering the columnNameTypeConstraint production.
	EnterColumnNameTypeConstraint(c *ColumnNameTypeConstraintContext)

	// EnterColumnConstraint is called when entering the columnConstraint production.
	EnterColumnConstraint(c *ColumnConstraintContext)

	// EnterForeignKeyConstraint is called when entering the foreignKeyConstraint production.
	EnterForeignKeyConstraint(c *ForeignKeyConstraintContext)

	// EnterColConstraint is called when entering the colConstraint production.
	EnterColConstraint(c *ColConstraintContext)

	// EnterAlterColumnConstraint is called when entering the alterColumnConstraint production.
	EnterAlterColumnConstraint(c *AlterColumnConstraintContext)

	// EnterAlterForeignKeyConstraint is called when entering the alterForeignKeyConstraint production.
	EnterAlterForeignKeyConstraint(c *AlterForeignKeyConstraintContext)

	// EnterAlterColConstraint is called when entering the alterColConstraint production.
	EnterAlterColConstraint(c *AlterColConstraintContext)

	// EnterColumnConstraintType is called when entering the columnConstraintType production.
	EnterColumnConstraintType(c *ColumnConstraintTypeContext)

	// EnterDefaultVal is called when entering the defaultVal production.
	EnterDefaultVal(c *DefaultValContext)

	// EnterTableConstraintType is called when entering the tableConstraintType production.
	EnterTableConstraintType(c *TableConstraintTypeContext)

	// EnterConstraintOptsCreate is called when entering the constraintOptsCreate production.
	EnterConstraintOptsCreate(c *ConstraintOptsCreateContext)

	// EnterConstraintOptsAlter is called when entering the constraintOptsAlter production.
	EnterConstraintOptsAlter(c *ConstraintOptsAlterContext)

	// EnterColumnNameColonType is called when entering the columnNameColonType production.
	EnterColumnNameColonType(c *ColumnNameColonTypeContext)

	// EnterColType is called when entering the colType production.
	EnterColType(c *ColTypeContext)

	// EnterColTypeList is called when entering the colTypeList production.
	EnterColTypeList(c *ColTypeListContext)

	// EnterType_db_col is called when entering the type_db_col production.
	EnterType_db_col(c *Type_db_colContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterListType is called when entering the listType production.
	EnterListType(c *ListTypeContext)

	// EnterStructType is called when entering the structType production.
	EnterStructType(c *StructTypeContext)

	// EnterMapType is called when entering the mapType production.
	EnterMapType(c *MapTypeContext)

	// EnterUnionType is called when entering the unionType production.
	EnterUnionType(c *UnionTypeContext)

	// EnterSetOperator is called when entering the setOperator production.
	EnterSetOperator(c *SetOperatorContext)

	// EnterQueryStatementExpression is called when entering the queryStatementExpression production.
	EnterQueryStatementExpression(c *QueryStatementExpressionContext)

	// EnterQueryStatementExpressionBody is called when entering the queryStatementExpressionBody production.
	EnterQueryStatementExpressionBody(c *QueryStatementExpressionBodyContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterCteStatement is called when entering the cteStatement production.
	EnterCteStatement(c *CteStatementContext)

	// EnterFromStatement is called when entering the fromStatement production.
	EnterFromStatement(c *FromStatementContext)

	// EnterSingleFromStatement is called when entering the singleFromStatement production.
	EnterSingleFromStatement(c *SingleFromStatementContext)

	// EnterRegularBody is called when entering the regularBody production.
	EnterRegularBody(c *RegularBodyContext)

	// EnterAtomSelectStatement is called when entering the atomSelectStatement production.
	EnterAtomSelectStatement(c *AtomSelectStatementContext)

	// EnterSelectStatement is called when entering the selectStatement production.
	EnterSelectStatement(c *SelectStatementContext)

	// EnterSetOpSelectStatement is called when entering the setOpSelectStatement production.
	EnterSetOpSelectStatement(c *SetOpSelectStatementContext)

	// EnterSelectStatementWithCTE is called when entering the selectStatementWithCTE production.
	EnterSelectStatementWithCTE(c *SelectStatementWithCTEContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterInsertClause is called when entering the insertClause production.
	EnterInsertClause(c *InsertClauseContext)

	// EnterDestination is called when entering the destination production.
	EnterDestination(c *DestinationContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterColumnAssignmentClause is called when entering the columnAssignmentClause production.
	EnterColumnAssignmentClause(c *ColumnAssignmentClauseContext)

	// EnterSetColumnsClause is called when entering the setColumnsClause production.
	EnterSetColumnsClause(c *SetColumnsClauseContext)

	// EnterUpdateStatement is called when entering the updateStatement production.
	EnterUpdateStatement(c *UpdateStatementContext)

	// EnterSqlTransactionStatement is called when entering the sqlTransactionStatement production.
	EnterSqlTransactionStatement(c *SqlTransactionStatementContext)

	// EnterStartTransactionStatement is called when entering the startTransactionStatement production.
	EnterStartTransactionStatement(c *StartTransactionStatementContext)

	// EnterTransactionMode is called when entering the transactionMode production.
	EnterTransactionMode(c *TransactionModeContext)

	// EnterTransactionAccessMode is called when entering the transactionAccessMode production.
	EnterTransactionAccessMode(c *TransactionAccessModeContext)

	// EnterIsolationLevel is called when entering the isolationLevel production.
	EnterIsolationLevel(c *IsolationLevelContext)

	// EnterLevelOfIsolation is called when entering the levelOfIsolation production.
	EnterLevelOfIsolation(c *LevelOfIsolationContext)

	// EnterCommitStatement is called when entering the commitStatement production.
	EnterCommitStatement(c *CommitStatementContext)

	// EnterRollbackStatement is called when entering the rollbackStatement production.
	EnterRollbackStatement(c *RollbackStatementContext)

	// EnterSetAutoCommitStatement is called when entering the setAutoCommitStatement production.
	EnterSetAutoCommitStatement(c *SetAutoCommitStatementContext)

	// EnterAbortTransactionStatement is called when entering the abortTransactionStatement production.
	EnterAbortTransactionStatement(c *AbortTransactionStatementContext)

	// EnterMergeStatement is called when entering the mergeStatement production.
	EnterMergeStatement(c *MergeStatementContext)

	// EnterWhenClauses is called when entering the whenClauses production.
	EnterWhenClauses(c *WhenClausesContext)

	// EnterWhenNotMatchedClause is called when entering the whenNotMatchedClause production.
	EnterWhenNotMatchedClause(c *WhenNotMatchedClauseContext)

	// EnterWhenMatchedAndClause is called when entering the whenMatchedAndClause production.
	EnterWhenMatchedAndClause(c *WhenMatchedAndClauseContext)

	// EnterWhenMatchedThenClause is called when entering the whenMatchedThenClause production.
	EnterWhenMatchedThenClause(c *WhenMatchedThenClauseContext)

	// EnterUpdateOrDelete is called when entering the updateOrDelete production.
	EnterUpdateOrDelete(c *UpdateOrDeleteContext)

	// EnterKillQueryStatement is called when entering the killQueryStatement production.
	EnterKillQueryStatement(c *KillQueryStatementContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterSelectList is called when entering the selectList production.
	EnterSelectList(c *SelectListContext)

	// EnterSelectTrfmClause is called when entering the selectTrfmClause production.
	EnterSelectTrfmClause(c *SelectTrfmClauseContext)

	// EnterSelectItem is called when entering the selectItem production.
	EnterSelectItem(c *SelectItemContext)

	// EnterTrfmClause is called when entering the trfmClause production.
	EnterTrfmClause(c *TrfmClauseContext)

	// EnterSelectExpression is called when entering the selectExpression production.
	EnterSelectExpression(c *SelectExpressionContext)

	// EnterSelectExpressionList is called when entering the selectExpressionList production.
	EnterSelectExpressionList(c *SelectExpressionListContext)

	// EnterWindow_clause is called when entering the window_clause production.
	EnterWindow_clause(c *Window_clauseContext)

	// EnterWindow_defn is called when entering the window_defn production.
	EnterWindow_defn(c *Window_defnContext)

	// EnterWindow_specification is called when entering the window_specification production.
	EnterWindow_specification(c *Window_specificationContext)

	// EnterWindow_frame is called when entering the window_frame production.
	EnterWindow_frame(c *Window_frameContext)

	// EnterWindow_range_expression is called when entering the window_range_expression production.
	EnterWindow_range_expression(c *Window_range_expressionContext)

	// EnterWindow_value_expression is called when entering the window_value_expression production.
	EnterWindow_value_expression(c *Window_value_expressionContext)

	// EnterWindow_frame_start_boundary is called when entering the window_frame_start_boundary production.
	EnterWindow_frame_start_boundary(c *Window_frame_start_boundaryContext)

	// EnterWindow_frame_boundary is called when entering the window_frame_boundary production.
	EnterWindow_frame_boundary(c *Window_frame_boundaryContext)

	// EnterTableAllColumns is called when entering the tableAllColumns production.
	EnterTableAllColumns(c *TableAllColumnsContext)

	// EnterTableOrColumn is called when entering the tableOrColumn production.
	EnterTableOrColumn(c *TableOrColumnContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterAliasList is called when entering the aliasList production.
	EnterAliasList(c *AliasListContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterFromSource is called when entering the fromSource production.
	EnterFromSource(c *FromSourceContext)

	// EnterAtomjoinSource is called when entering the atomjoinSource production.
	EnterAtomjoinSource(c *AtomjoinSourceContext)

	// EnterJoinSource is called when entering the joinSource production.
	EnterJoinSource(c *JoinSourceContext)

	// EnterJoinSourcePart is called when entering the joinSourcePart production.
	EnterJoinSourcePart(c *JoinSourcePartContext)

	// EnterUniqueJoinSource is called when entering the uniqueJoinSource production.
	EnterUniqueJoinSource(c *UniqueJoinSourceContext)

	// EnterUniqueJoinExpr is called when entering the uniqueJoinExpr production.
	EnterUniqueJoinExpr(c *UniqueJoinExprContext)

	// EnterUniqueJoinToken is called when entering the uniqueJoinToken production.
	EnterUniqueJoinToken(c *UniqueJoinTokenContext)

	// EnterJoinToken is called when entering the joinToken production.
	EnterJoinToken(c *JoinTokenContext)

	// EnterLateralView is called when entering the lateralView production.
	EnterLateralView(c *LateralViewContext)

	// EnterTableAlias is called when entering the tableAlias production.
	EnterTableAlias(c *TableAliasContext)

	// EnterTableBucketSample is called when entering the tableBucketSample production.
	EnterTableBucketSample(c *TableBucketSampleContext)

	// EnterSplitSample is called when entering the splitSample production.
	EnterSplitSample(c *SplitSampleContext)

	// EnterTableSample is called when entering the tableSample production.
	EnterTableSample(c *TableSampleContext)

	// EnterTableSource is called when entering the tableSource production.
	EnterTableSource(c *TableSourceContext)

	// EnterUniqueJoinTableSource is called when entering the uniqueJoinTableSource production.
	EnterUniqueJoinTableSource(c *UniqueJoinTableSourceContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterViewName is called when entering the viewName production.
	EnterViewName(c *ViewNameContext)

	// EnterSubQuerySource is called when entering the subQuerySource production.
	EnterSubQuerySource(c *SubQuerySourceContext)

	// EnterPartitioningSpec is called when entering the partitioningSpec production.
	EnterPartitioningSpec(c *PartitioningSpecContext)

	// EnterPartitionTableFunctionSource is called when entering the partitionTableFunctionSource production.
	EnterPartitionTableFunctionSource(c *PartitionTableFunctionSourceContext)

	// EnterPartitionedTableFunction is called when entering the partitionedTableFunction production.
	EnterPartitionedTableFunction(c *PartitionedTableFunctionContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterSearchCondition is called when entering the searchCondition production.
	EnterSearchCondition(c *SearchConditionContext)

	// EnterValuesClause is called when entering the valuesClause production.
	EnterValuesClause(c *ValuesClauseContext)

	// EnterValuesTableConstructor is called when entering the valuesTableConstructor production.
	EnterValuesTableConstructor(c *ValuesTableConstructorContext)

	// EnterValueRowConstructor is called when entering the valueRowConstructor production.
	EnterValueRowConstructor(c *ValueRowConstructorContext)

	// EnterVirtualTableSource is called when entering the virtualTableSource production.
	EnterVirtualTableSource(c *VirtualTableSourceContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterGroupby_expression is called when entering the groupby_expression production.
	EnterGroupby_expression(c *Groupby_expressionContext)

	// EnterGroupByEmpty is called when entering the groupByEmpty production.
	EnterGroupByEmpty(c *GroupByEmptyContext)

	// EnterRollupStandard is called when entering the rollupStandard production.
	EnterRollupStandard(c *RollupStandardContext)

	// EnterRollupOldSyntax is called when entering the rollupOldSyntax production.
	EnterRollupOldSyntax(c *RollupOldSyntaxContext)

	// EnterGroupingSetExpression is called when entering the groupingSetExpression production.
	EnterGroupingSetExpression(c *GroupingSetExpressionContext)

	// EnterGroupingSetExpressionMultiple is called when entering the groupingSetExpressionMultiple production.
	EnterGroupingSetExpressionMultiple(c *GroupingSetExpressionMultipleContext)

	// EnterGroupingExpressionSingle is called when entering the groupingExpressionSingle production.
	EnterGroupingExpressionSingle(c *GroupingExpressionSingleContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterHavingCondition is called when entering the havingCondition production.
	EnterHavingCondition(c *HavingConditionContext)

	// EnterExpressionsInParenthesis is called when entering the expressionsInParenthesis production.
	EnterExpressionsInParenthesis(c *ExpressionsInParenthesisContext)

	// EnterExpressionsNotInParenthesis is called when entering the expressionsNotInParenthesis production.
	EnterExpressionsNotInParenthesis(c *ExpressionsNotInParenthesisContext)

	// EnterExpressionPart is called when entering the expressionPart production.
	EnterExpressionPart(c *ExpressionPartContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterColumnRefOrderInParenthesis is called when entering the columnRefOrderInParenthesis production.
	EnterColumnRefOrderInParenthesis(c *ColumnRefOrderInParenthesisContext)

	// EnterColumnRefOrderNotInParenthesis is called when entering the columnRefOrderNotInParenthesis production.
	EnterColumnRefOrderNotInParenthesis(c *ColumnRefOrderNotInParenthesisContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterClusterByClause is called when entering the clusterByClause production.
	EnterClusterByClause(c *ClusterByClauseContext)

	// EnterPartitionByClause is called when entering the partitionByClause production.
	EnterPartitionByClause(c *PartitionByClauseContext)

	// EnterDistributeByClause is called when entering the distributeByClause production.
	EnterDistributeByClause(c *DistributeByClauseContext)

	// EnterSortByClause is called when entering the sortByClause production.
	EnterSortByClause(c *SortByClauseContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterCastExpression is called when entering the castExpression production.
	EnterCastExpression(c *CastExpressionContext)

	// EnterCaseExpression is called when entering the caseExpression production.
	EnterCaseExpression(c *CaseExpressionContext)

	// EnterWhenExpression is called when entering the whenExpression production.
	EnterWhenExpression(c *WhenExpressionContext)

	// EnterFloorExpression is called when entering the floorExpression production.
	EnterFloorExpression(c *FloorExpressionContext)

	// EnterFloorDateQualifiers is called when entering the floorDateQualifiers production.
	EnterFloorDateQualifiers(c *FloorDateQualifiersContext)

	// EnterExtractExpression is called when entering the extractExpression production.
	EnterExtractExpression(c *ExtractExpressionContext)

	// EnterTimeQualifiers is called when entering the timeQualifiers production.
	EnterTimeQualifiers(c *TimeQualifiersContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterStringLiteralSequence is called when entering the stringLiteralSequence production.
	EnterStringLiteralSequence(c *StringLiteralSequenceContext)

	// EnterCharSetStringLiteral is called when entering the charSetStringLiteral production.
	EnterCharSetStringLiteral(c *CharSetStringLiteralContext)

	// EnterDateLiteral is called when entering the dateLiteral production.
	EnterDateLiteral(c *DateLiteralContext)

	// EnterTimestampLiteral is called when entering the timestampLiteral production.
	EnterTimestampLiteral(c *TimestampLiteralContext)

	// EnterTimestampLocalTZLiteral is called when entering the timestampLocalTZLiteral production.
	EnterTimestampLocalTZLiteral(c *TimestampLocalTZLiteralContext)

	// EnterIntervalValue is called when entering the intervalValue production.
	EnterIntervalValue(c *IntervalValueContext)

	// EnterIntervalLiteral is called when entering the intervalLiteral production.
	EnterIntervalLiteral(c *IntervalLiteralContext)

	// EnterIntervalExpression is called when entering the intervalExpression production.
	EnterIntervalExpression(c *IntervalExpressionContext)

	// EnterIntervalQualifiers is called when entering the intervalQualifiers production.
	EnterIntervalQualifiers(c *IntervalQualifiersContext)

	// EnterAtomExpression is called when entering the atomExpression production.
	EnterAtomExpression(c *AtomExpressionContext)

	// EnterPrecedenceUnaryOperator is called when entering the precedenceUnaryOperator production.
	EnterPrecedenceUnaryOperator(c *PrecedenceUnaryOperatorContext)

	// EnterIsCondition is called when entering the isCondition production.
	EnterIsCondition(c *IsConditionContext)

	// EnterPrecedenceBitwiseXorOperator is called when entering the precedenceBitwiseXorOperator production.
	EnterPrecedenceBitwiseXorOperator(c *PrecedenceBitwiseXorOperatorContext)

	// EnterPrecedenceStarOperator is called when entering the precedenceStarOperator production.
	EnterPrecedenceStarOperator(c *PrecedenceStarOperatorContext)

	// EnterPrecedencePlusOperator is called when entering the precedencePlusOperator production.
	EnterPrecedencePlusOperator(c *PrecedencePlusOperatorContext)

	// EnterPrecedenceConcatenateOperator is called when entering the precedenceConcatenateOperator production.
	EnterPrecedenceConcatenateOperator(c *PrecedenceConcatenateOperatorContext)

	// EnterPrecedenceAmpersandOperator is called when entering the precedenceAmpersandOperator production.
	EnterPrecedenceAmpersandOperator(c *PrecedenceAmpersandOperatorContext)

	// EnterPrecedenceBitwiseOrOperator is called when entering the precedenceBitwiseOrOperator production.
	EnterPrecedenceBitwiseOrOperator(c *PrecedenceBitwiseOrOperatorContext)

	// EnterPrecedenceRegexpOperator is called when entering the precedenceRegexpOperator production.
	EnterPrecedenceRegexpOperator(c *PrecedenceRegexpOperatorContext)

	// EnterPrecedenceSimilarOperator is called when entering the precedenceSimilarOperator production.
	EnterPrecedenceSimilarOperator(c *PrecedenceSimilarOperatorContext)

	// EnterPrecedenceDistinctOperator is called when entering the precedenceDistinctOperator production.
	EnterPrecedenceDistinctOperator(c *PrecedenceDistinctOperatorContext)

	// EnterPrecedenceEqualOperator is called when entering the precedenceEqualOperator production.
	EnterPrecedenceEqualOperator(c *PrecedenceEqualOperatorContext)

	// EnterPrecedenceNotOperator is called when entering the precedenceNotOperator production.
	EnterPrecedenceNotOperator(c *PrecedenceNotOperatorContext)

	// EnterPrecedenceAndOperator is called when entering the precedenceAndOperator production.
	EnterPrecedenceAndOperator(c *PrecedenceAndOperatorContext)

	// EnterPrecedenceOrOperator is called when entering the precedenceOrOperator production.
	EnterPrecedenceOrOperator(c *PrecedenceOrOperatorContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSubQueryExpression is called when entering the subQueryExpression production.
	EnterSubQueryExpression(c *SubQueryExpressionContext)

	// EnterPrecedenceSimilarExpressionPart is called when entering the precedenceSimilarExpressionPart production.
	EnterPrecedenceSimilarExpressionPart(c *PrecedenceSimilarExpressionPartContext)

	// EnterPrecedenceSimilarExpressionAtom is called when entering the precedenceSimilarExpressionAtom production.
	EnterPrecedenceSimilarExpressionAtom(c *PrecedenceSimilarExpressionAtomContext)

	// EnterPrecedenceSimilarExpressionIn is called when entering the precedenceSimilarExpressionIn production.
	EnterPrecedenceSimilarExpressionIn(c *PrecedenceSimilarExpressionInContext)

	// EnterPrecedenceSimilarExpressionPartNot is called when entering the precedenceSimilarExpressionPartNot production.
	EnterPrecedenceSimilarExpressionPartNot(c *PrecedenceSimilarExpressionPartNotContext)

	// EnterBooleanValue is called when entering the booleanValue production.
	EnterBooleanValue(c *BooleanValueContext)

	// EnterBooleanValueTok is called when entering the booleanValueTok production.
	EnterBooleanValueTok(c *BooleanValueTokContext)

	// EnterTableOrPartition is called when entering the tableOrPartition production.
	EnterTableOrPartition(c *TableOrPartitionContext)

	// EnterPartitionSpec is called when entering the partitionSpec production.
	EnterPartitionSpec(c *PartitionSpecContext)

	// EnterPartitionVal is called when entering the partitionVal production.
	EnterPartitionVal(c *PartitionValContext)

	// EnterDropPartitionSpec is called when entering the dropPartitionSpec production.
	EnterDropPartitionSpec(c *DropPartitionSpecContext)

	// EnterDropPartitionVal is called when entering the dropPartitionVal production.
	EnterDropPartitionVal(c *DropPartitionValContext)

	// EnterDropPartitionOperator is called when entering the dropPartitionOperator production.
	EnterDropPartitionOperator(c *DropPartitionOperatorContext)

	// EnterSysFuncNames is called when entering the sysFuncNames production.
	EnterSysFuncNames(c *SysFuncNamesContext)

	// EnterDescFuncNames is called when entering the descFuncNames production.
	EnterDescFuncNames(c *DescFuncNamesContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterFunctionIdentifier is called when entering the functionIdentifier production.
	EnterFunctionIdentifier(c *FunctionIdentifierContext)

	// EnterPrincipalIdentifier is called when entering the principalIdentifier production.
	EnterPrincipalIdentifier(c *PrincipalIdentifierContext)

	// EnterNonReserved is called when entering the nonReserved production.
	EnterNonReserved(c *NonReservedContext)

	// EnterSql11ReservedKeywordsUsedAsFunctionName is called when entering the sql11ReservedKeywordsUsedAsFunctionName production.
	EnterSql11ReservedKeywordsUsedAsFunctionName(c *Sql11ReservedKeywordsUsedAsFunctionNameContext)

	// EnterResourcePlanDdlStatements is called when entering the resourcePlanDdlStatements production.
	EnterResourcePlanDdlStatements(c *ResourcePlanDdlStatementsContext)

	// EnterRpAssign is called when entering the rpAssign production.
	EnterRpAssign(c *RpAssignContext)

	// EnterRpAssignList is called when entering the rpAssignList production.
	EnterRpAssignList(c *RpAssignListContext)

	// EnterRpUnassign is called when entering the rpUnassign production.
	EnterRpUnassign(c *RpUnassignContext)

	// EnterRpUnassignList is called when entering the rpUnassignList production.
	EnterRpUnassignList(c *RpUnassignListContext)

	// EnterCreateResourcePlanStatement is called when entering the createResourcePlanStatement production.
	EnterCreateResourcePlanStatement(c *CreateResourcePlanStatementContext)

	// EnterWithReplace is called when entering the withReplace production.
	EnterWithReplace(c *WithReplaceContext)

	// EnterActivate is called when entering the activate production.
	EnterActivate(c *ActivateContext)

	// EnterEnable is called when entering the enable production.
	EnterEnable(c *EnableContext)

	// EnterDisable is called when entering the disable production.
	EnterDisable(c *DisableContext)

	// EnterUnmanaged is called when entering the unmanaged production.
	EnterUnmanaged(c *UnmanagedContext)

	// EnterAlterResourcePlanStatement is called when entering the alterResourcePlanStatement production.
	EnterAlterResourcePlanStatement(c *AlterResourcePlanStatementContext)

	// EnterGlobalWmStatement is called when entering the globalWmStatement production.
	EnterGlobalWmStatement(c *GlobalWmStatementContext)

	// EnterReplaceResourcePlanStatement is called when entering the replaceResourcePlanStatement production.
	EnterReplaceResourcePlanStatement(c *ReplaceResourcePlanStatementContext)

	// EnterDropResourcePlanStatement is called when entering the dropResourcePlanStatement production.
	EnterDropResourcePlanStatement(c *DropResourcePlanStatementContext)

	// EnterPoolPath is called when entering the poolPath production.
	EnterPoolPath(c *PoolPathContext)

	// EnterTriggerExpression is called when entering the triggerExpression production.
	EnterTriggerExpression(c *TriggerExpressionContext)

	// EnterTriggerExpressionStandalone is called when entering the triggerExpressionStandalone production.
	EnterTriggerExpressionStandalone(c *TriggerExpressionStandaloneContext)

	// EnterTriggerOrExpression is called when entering the triggerOrExpression production.
	EnterTriggerOrExpression(c *TriggerOrExpressionContext)

	// EnterTriggerAndExpression is called when entering the triggerAndExpression production.
	EnterTriggerAndExpression(c *TriggerAndExpressionContext)

	// EnterTriggerAtomExpression is called when entering the triggerAtomExpression production.
	EnterTriggerAtomExpression(c *TriggerAtomExpressionContext)

	// EnterTriggerLiteral is called when entering the triggerLiteral production.
	EnterTriggerLiteral(c *TriggerLiteralContext)

	// EnterComparisionOperator is called when entering the comparisionOperator production.
	EnterComparisionOperator(c *ComparisionOperatorContext)

	// EnterTriggerActionExpression is called when entering the triggerActionExpression production.
	EnterTriggerActionExpression(c *TriggerActionExpressionContext)

	// EnterTriggerActionExpressionStandalone is called when entering the triggerActionExpressionStandalone production.
	EnterTriggerActionExpressionStandalone(c *TriggerActionExpressionStandaloneContext)

	// EnterCreateTriggerStatement is called when entering the createTriggerStatement production.
	EnterCreateTriggerStatement(c *CreateTriggerStatementContext)

	// EnterAlterTriggerStatement is called when entering the alterTriggerStatement production.
	EnterAlterTriggerStatement(c *AlterTriggerStatementContext)

	// EnterDropTriggerStatement is called when entering the dropTriggerStatement production.
	EnterDropTriggerStatement(c *DropTriggerStatementContext)

	// EnterPoolAssign is called when entering the poolAssign production.
	EnterPoolAssign(c *PoolAssignContext)

	// EnterPoolAssignList is called when entering the poolAssignList production.
	EnterPoolAssignList(c *PoolAssignListContext)

	// EnterCreatePoolStatement is called when entering the createPoolStatement production.
	EnterCreatePoolStatement(c *CreatePoolStatementContext)

	// EnterAlterPoolStatement is called when entering the alterPoolStatement production.
	EnterAlterPoolStatement(c *AlterPoolStatementContext)

	// EnterDropPoolStatement is called when entering the dropPoolStatement production.
	EnterDropPoolStatement(c *DropPoolStatementContext)

	// EnterCreateMappingStatement is called when entering the createMappingStatement production.
	EnterCreateMappingStatement(c *CreateMappingStatementContext)

	// EnterAlterMappingStatement is called when entering the alterMappingStatement production.
	EnterAlterMappingStatement(c *AlterMappingStatementContext)

	// EnterDropMappingStatement is called when entering the dropMappingStatement production.
	EnterDropMappingStatement(c *DropMappingStatementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitExplainStatement is called when exiting the explainStatement production.
	ExitExplainStatement(c *ExplainStatementContext)

	// ExitExplainOption is called when exiting the explainOption production.
	ExitExplainOption(c *ExplainOptionContext)

	// ExitVectorizationOnly is called when exiting the vectorizationOnly production.
	ExitVectorizationOnly(c *VectorizationOnlyContext)

	// ExitVectorizatonDetail is called when exiting the vectorizatonDetail production.
	ExitVectorizatonDetail(c *VectorizatonDetailContext)

	// ExitExecStatement is called when exiting the execStatement production.
	ExitExecStatement(c *ExecStatementContext)

	// ExitLoadStatement is called when exiting the loadStatement production.
	ExitLoadStatement(c *LoadStatementContext)

	// ExitReplicationClause is called when exiting the replicationClause production.
	ExitReplicationClause(c *ReplicationClauseContext)

	// ExitExportStatement is called when exiting the exportStatement production.
	ExitExportStatement(c *ExportStatementContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitReplDumpStatement is called when exiting the replDumpStatement production.
	ExitReplDumpStatement(c *ReplDumpStatementContext)

	// ExitReplLoadStatement is called when exiting the replLoadStatement production.
	ExitReplLoadStatement(c *ReplLoadStatementContext)

	// ExitReplConfigs is called when exiting the replConfigs production.
	ExitReplConfigs(c *ReplConfigsContext)

	// ExitReplConfigsList is called when exiting the replConfigsList production.
	ExitReplConfigsList(c *ReplConfigsListContext)

	// ExitReplStatusStatement is called when exiting the replStatusStatement production.
	ExitReplStatusStatement(c *ReplStatusStatementContext)

	// ExitDdlStatement is called when exiting the ddlStatement production.
	ExitDdlStatement(c *DdlStatementContext)

	// ExitIfExists is called when exiting the ifExists production.
	ExitIfExists(c *IfExistsContext)

	// ExitRestrictOrCascade is called when exiting the restrictOrCascade production.
	ExitRestrictOrCascade(c *RestrictOrCascadeContext)

	// ExitIfNotExists is called when exiting the ifNotExists production.
	ExitIfNotExists(c *IfNotExistsContext)

	// ExitRewriteEnabled is called when exiting the rewriteEnabled production.
	ExitRewriteEnabled(c *RewriteEnabledContext)

	// ExitRewriteDisabled is called when exiting the rewriteDisabled production.
	ExitRewriteDisabled(c *RewriteDisabledContext)

	// ExitStoredAsDirs is called when exiting the storedAsDirs production.
	ExitStoredAsDirs(c *StoredAsDirsContext)

	// ExitOrReplace is called when exiting the orReplace production.
	ExitOrReplace(c *OrReplaceContext)

	// ExitCreateDatabaseStatement is called when exiting the createDatabaseStatement production.
	ExitCreateDatabaseStatement(c *CreateDatabaseStatementContext)

	// ExitDbLocation is called when exiting the dbLocation production.
	ExitDbLocation(c *DbLocationContext)

	// ExitDbProperties is called when exiting the dbProperties production.
	ExitDbProperties(c *DbPropertiesContext)

	// ExitDbPropertiesList is called when exiting the dbPropertiesList production.
	ExitDbPropertiesList(c *DbPropertiesListContext)

	// ExitSwitchDatabaseStatement is called when exiting the switchDatabaseStatement production.
	ExitSwitchDatabaseStatement(c *SwitchDatabaseStatementContext)

	// ExitDropDatabaseStatement is called when exiting the dropDatabaseStatement production.
	ExitDropDatabaseStatement(c *DropDatabaseStatementContext)

	// ExitDatabaseComment is called when exiting the databaseComment production.
	ExitDatabaseComment(c *DatabaseCommentContext)

	// ExitCreateTableStatement is called when exiting the createTableStatement production.
	ExitCreateTableStatement(c *CreateTableStatementContext)

	// ExitTruncateTableStatement is called when exiting the truncateTableStatement production.
	ExitTruncateTableStatement(c *TruncateTableStatementContext)

	// ExitDropTableStatement is called when exiting the dropTableStatement production.
	ExitDropTableStatement(c *DropTableStatementContext)

	// ExitAlterStatement is called when exiting the alterStatement production.
	ExitAlterStatement(c *AlterStatementContext)

	// ExitAlterTableStatementSuffix is called when exiting the alterTableStatementSuffix production.
	ExitAlterTableStatementSuffix(c *AlterTableStatementSuffixContext)

	// ExitAlterTblPartitionStatementSuffix is called when exiting the alterTblPartitionStatementSuffix production.
	ExitAlterTblPartitionStatementSuffix(c *AlterTblPartitionStatementSuffixContext)

	// ExitAlterStatementPartitionKeyType is called when exiting the alterStatementPartitionKeyType production.
	ExitAlterStatementPartitionKeyType(c *AlterStatementPartitionKeyTypeContext)

	// ExitAlterViewStatementSuffix is called when exiting the alterViewStatementSuffix production.
	ExitAlterViewStatementSuffix(c *AlterViewStatementSuffixContext)

	// ExitAlterMaterializedViewStatementSuffix is called when exiting the alterMaterializedViewStatementSuffix production.
	ExitAlterMaterializedViewStatementSuffix(c *AlterMaterializedViewStatementSuffixContext)

	// ExitAlterDatabaseStatementSuffix is called when exiting the alterDatabaseStatementSuffix production.
	ExitAlterDatabaseStatementSuffix(c *AlterDatabaseStatementSuffixContext)

	// ExitAlterDatabaseSuffixProperties is called when exiting the alterDatabaseSuffixProperties production.
	ExitAlterDatabaseSuffixProperties(c *AlterDatabaseSuffixPropertiesContext)

	// ExitAlterDatabaseSuffixSetOwner is called when exiting the alterDatabaseSuffixSetOwner production.
	ExitAlterDatabaseSuffixSetOwner(c *AlterDatabaseSuffixSetOwnerContext)

	// ExitAlterDatabaseSuffixSetLocation is called when exiting the alterDatabaseSuffixSetLocation production.
	ExitAlterDatabaseSuffixSetLocation(c *AlterDatabaseSuffixSetLocationContext)

	// ExitAlterStatementSuffixRename is called when exiting the alterStatementSuffixRename production.
	ExitAlterStatementSuffixRename(c *AlterStatementSuffixRenameContext)

	// ExitAlterStatementSuffixAddCol is called when exiting the alterStatementSuffixAddCol production.
	ExitAlterStatementSuffixAddCol(c *AlterStatementSuffixAddColContext)

	// ExitAlterStatementSuffixAddConstraint is called when exiting the alterStatementSuffixAddConstraint production.
	ExitAlterStatementSuffixAddConstraint(c *AlterStatementSuffixAddConstraintContext)

	// ExitAlterStatementSuffixUpdateColumns is called when exiting the alterStatementSuffixUpdateColumns production.
	ExitAlterStatementSuffixUpdateColumns(c *AlterStatementSuffixUpdateColumnsContext)

	// ExitAlterStatementSuffixDropConstraint is called when exiting the alterStatementSuffixDropConstraint production.
	ExitAlterStatementSuffixDropConstraint(c *AlterStatementSuffixDropConstraintContext)

	// ExitAlterStatementSuffixRenameCol is called when exiting the alterStatementSuffixRenameCol production.
	ExitAlterStatementSuffixRenameCol(c *AlterStatementSuffixRenameColContext)

	// ExitAlterStatementSuffixUpdateStatsCol is called when exiting the alterStatementSuffixUpdateStatsCol production.
	ExitAlterStatementSuffixUpdateStatsCol(c *AlterStatementSuffixUpdateStatsColContext)

	// ExitAlterStatementSuffixUpdateStats is called when exiting the alterStatementSuffixUpdateStats production.
	ExitAlterStatementSuffixUpdateStats(c *AlterStatementSuffixUpdateStatsContext)

	// ExitAlterStatementChangeColPosition is called when exiting the alterStatementChangeColPosition production.
	ExitAlterStatementChangeColPosition(c *AlterStatementChangeColPositionContext)

	// ExitAlterStatementSuffixAddPartitions is called when exiting the alterStatementSuffixAddPartitions production.
	ExitAlterStatementSuffixAddPartitions(c *AlterStatementSuffixAddPartitionsContext)

	// ExitAlterStatementSuffixAddPartitionsElement is called when exiting the alterStatementSuffixAddPartitionsElement production.
	ExitAlterStatementSuffixAddPartitionsElement(c *AlterStatementSuffixAddPartitionsElementContext)

	// ExitAlterStatementSuffixTouch is called when exiting the alterStatementSuffixTouch production.
	ExitAlterStatementSuffixTouch(c *AlterStatementSuffixTouchContext)

	// ExitAlterStatementSuffixArchive is called when exiting the alterStatementSuffixArchive production.
	ExitAlterStatementSuffixArchive(c *AlterStatementSuffixArchiveContext)

	// ExitAlterStatementSuffixUnArchive is called when exiting the alterStatementSuffixUnArchive production.
	ExitAlterStatementSuffixUnArchive(c *AlterStatementSuffixUnArchiveContext)

	// ExitPartitionLocation is called when exiting the partitionLocation production.
	ExitPartitionLocation(c *PartitionLocationContext)

	// ExitAlterStatementSuffixDropPartitions is called when exiting the alterStatementSuffixDropPartitions production.
	ExitAlterStatementSuffixDropPartitions(c *AlterStatementSuffixDropPartitionsContext)

	// ExitAlterStatementSuffixProperties is called when exiting the alterStatementSuffixProperties production.
	ExitAlterStatementSuffixProperties(c *AlterStatementSuffixPropertiesContext)

	// ExitAlterViewSuffixProperties is called when exiting the alterViewSuffixProperties production.
	ExitAlterViewSuffixProperties(c *AlterViewSuffixPropertiesContext)

	// ExitAlterMaterializedViewSuffixRewrite is called when exiting the alterMaterializedViewSuffixRewrite production.
	ExitAlterMaterializedViewSuffixRewrite(c *AlterMaterializedViewSuffixRewriteContext)

	// ExitAlterMaterializedViewSuffixRebuild is called when exiting the alterMaterializedViewSuffixRebuild production.
	ExitAlterMaterializedViewSuffixRebuild(c *AlterMaterializedViewSuffixRebuildContext)

	// ExitAlterStatementSuffixSerdeProperties is called when exiting the alterStatementSuffixSerdeProperties production.
	ExitAlterStatementSuffixSerdeProperties(c *AlterStatementSuffixSerdePropertiesContext)

	// ExitTablePartitionPrefix is called when exiting the tablePartitionPrefix production.
	ExitTablePartitionPrefix(c *TablePartitionPrefixContext)

	// ExitAlterStatementSuffixFileFormat is called when exiting the alterStatementSuffixFileFormat production.
	ExitAlterStatementSuffixFileFormat(c *AlterStatementSuffixFileFormatContext)

	// ExitAlterStatementSuffixClusterbySortby is called when exiting the alterStatementSuffixClusterbySortby production.
	ExitAlterStatementSuffixClusterbySortby(c *AlterStatementSuffixClusterbySortbyContext)

	// ExitAlterTblPartitionStatementSuffixSkewedLocation is called when exiting the alterTblPartitionStatementSuffixSkewedLocation production.
	ExitAlterTblPartitionStatementSuffixSkewedLocation(c *AlterTblPartitionStatementSuffixSkewedLocationContext)

	// ExitSkewedLocations is called when exiting the skewedLocations production.
	ExitSkewedLocations(c *SkewedLocationsContext)

	// ExitSkewedLocationsList is called when exiting the skewedLocationsList production.
	ExitSkewedLocationsList(c *SkewedLocationsListContext)

	// ExitSkewedLocationMap is called when exiting the skewedLocationMap production.
	ExitSkewedLocationMap(c *SkewedLocationMapContext)

	// ExitAlterStatementSuffixLocation is called when exiting the alterStatementSuffixLocation production.
	ExitAlterStatementSuffixLocation(c *AlterStatementSuffixLocationContext)

	// ExitAlterStatementSuffixSkewedby is called when exiting the alterStatementSuffixSkewedby production.
	ExitAlterStatementSuffixSkewedby(c *AlterStatementSuffixSkewedbyContext)

	// ExitAlterStatementSuffixExchangePartition is called when exiting the alterStatementSuffixExchangePartition production.
	ExitAlterStatementSuffixExchangePartition(c *AlterStatementSuffixExchangePartitionContext)

	// ExitAlterStatementSuffixRenamePart is called when exiting the alterStatementSuffixRenamePart production.
	ExitAlterStatementSuffixRenamePart(c *AlterStatementSuffixRenamePartContext)

	// ExitAlterStatementSuffixStatsPart is called when exiting the alterStatementSuffixStatsPart production.
	ExitAlterStatementSuffixStatsPart(c *AlterStatementSuffixStatsPartContext)

	// ExitAlterStatementSuffixMergeFiles is called when exiting the alterStatementSuffixMergeFiles production.
	ExitAlterStatementSuffixMergeFiles(c *AlterStatementSuffixMergeFilesContext)

	// ExitAlterStatementSuffixBucketNum is called when exiting the alterStatementSuffixBucketNum production.
	ExitAlterStatementSuffixBucketNum(c *AlterStatementSuffixBucketNumContext)

	// ExitBlocking is called when exiting the blocking production.
	ExitBlocking(c *BlockingContext)

	// ExitAlterStatementSuffixCompact is called when exiting the alterStatementSuffixCompact production.
	ExitAlterStatementSuffixCompact(c *AlterStatementSuffixCompactContext)

	// ExitAlterStatementSuffixSetOwner is called when exiting the alterStatementSuffixSetOwner production.
	ExitAlterStatementSuffixSetOwner(c *AlterStatementSuffixSetOwnerContext)

	// ExitFileFormat is called when exiting the fileFormat production.
	ExitFileFormat(c *FileFormatContext)

	// ExitInputFileFormat is called when exiting the inputFileFormat production.
	ExitInputFileFormat(c *InputFileFormatContext)

	// ExitTabTypeExpr is called when exiting the tabTypeExpr production.
	ExitTabTypeExpr(c *TabTypeExprContext)

	// ExitPartTypeExpr is called when exiting the partTypeExpr production.
	ExitPartTypeExpr(c *PartTypeExprContext)

	// ExitTabPartColTypeExpr is called when exiting the tabPartColTypeExpr production.
	ExitTabPartColTypeExpr(c *TabPartColTypeExprContext)

	// ExitDescStatement is called when exiting the descStatement production.
	ExitDescStatement(c *DescStatementContext)

	// ExitAnalyzeStatement is called when exiting the analyzeStatement production.
	ExitAnalyzeStatement(c *AnalyzeStatementContext)

	// ExitShowStatement is called when exiting the showStatement production.
	ExitShowStatement(c *ShowStatementContext)

	// ExitLockStatement is called when exiting the lockStatement production.
	ExitLockStatement(c *LockStatementContext)

	// ExitLockDatabase is called when exiting the lockDatabase production.
	ExitLockDatabase(c *LockDatabaseContext)

	// ExitLockMode is called when exiting the lockMode production.
	ExitLockMode(c *LockModeContext)

	// ExitUnlockStatement is called when exiting the unlockStatement production.
	ExitUnlockStatement(c *UnlockStatementContext)

	// ExitUnlockDatabase is called when exiting the unlockDatabase production.
	ExitUnlockDatabase(c *UnlockDatabaseContext)

	// ExitCreateRoleStatement is called when exiting the createRoleStatement production.
	ExitCreateRoleStatement(c *CreateRoleStatementContext)

	// ExitDropRoleStatement is called when exiting the dropRoleStatement production.
	ExitDropRoleStatement(c *DropRoleStatementContext)

	// ExitGrantPrivileges is called when exiting the grantPrivileges production.
	ExitGrantPrivileges(c *GrantPrivilegesContext)

	// ExitRevokePrivileges is called when exiting the revokePrivileges production.
	ExitRevokePrivileges(c *RevokePrivilegesContext)

	// ExitGrantRole is called when exiting the grantRole production.
	ExitGrantRole(c *GrantRoleContext)

	// ExitRevokeRole is called when exiting the revokeRole production.
	ExitRevokeRole(c *RevokeRoleContext)

	// ExitShowRoleGrants is called when exiting the showRoleGrants production.
	ExitShowRoleGrants(c *ShowRoleGrantsContext)

	// ExitShowRoles is called when exiting the showRoles production.
	ExitShowRoles(c *ShowRolesContext)

	// ExitShowCurrentRole is called when exiting the showCurrentRole production.
	ExitShowCurrentRole(c *ShowCurrentRoleContext)

	// ExitSetRole is called when exiting the setRole production.
	ExitSetRole(c *SetRoleContext)

	// ExitShowGrants is called when exiting the showGrants production.
	ExitShowGrants(c *ShowGrantsContext)

	// ExitShowRolePrincipals is called when exiting the showRolePrincipals production.
	ExitShowRolePrincipals(c *ShowRolePrincipalsContext)

	// ExitPrivilegeIncludeColObject is called when exiting the privilegeIncludeColObject production.
	ExitPrivilegeIncludeColObject(c *PrivilegeIncludeColObjectContext)

	// ExitPrivilegeObject is called when exiting the privilegeObject production.
	ExitPrivilegeObject(c *PrivilegeObjectContext)

	// ExitPrivObject is called when exiting the privObject production.
	ExitPrivObject(c *PrivObjectContext)

	// ExitPrivObjectCols is called when exiting the privObjectCols production.
	ExitPrivObjectCols(c *PrivObjectColsContext)

	// ExitPrivilegeList is called when exiting the privilegeList production.
	ExitPrivilegeList(c *PrivilegeListContext)

	// ExitPrivlegeDef is called when exiting the privlegeDef production.
	ExitPrivlegeDef(c *PrivlegeDefContext)

	// ExitPrivilegeType is called when exiting the privilegeType production.
	ExitPrivilegeType(c *PrivilegeTypeContext)

	// ExitPrincipalSpecification is called when exiting the principalSpecification production.
	ExitPrincipalSpecification(c *PrincipalSpecificationContext)

	// ExitPrincipalName is called when exiting the principalName production.
	ExitPrincipalName(c *PrincipalNameContext)

	// ExitWithGrantOption is called when exiting the withGrantOption production.
	ExitWithGrantOption(c *WithGrantOptionContext)

	// ExitGrantOptionFor is called when exiting the grantOptionFor production.
	ExitGrantOptionFor(c *GrantOptionForContext)

	// ExitAdminOptionFor is called when exiting the adminOptionFor production.
	ExitAdminOptionFor(c *AdminOptionForContext)

	// ExitWithAdminOption is called when exiting the withAdminOption production.
	ExitWithAdminOption(c *WithAdminOptionContext)

	// ExitMetastoreCheck is called when exiting the metastoreCheck production.
	ExitMetastoreCheck(c *MetastoreCheckContext)

	// ExitResourceList is called when exiting the resourceList production.
	ExitResourceList(c *ResourceListContext)

	// ExitResource is called when exiting the resource production.
	ExitResource(c *ResourceContext)

	// ExitResourceType is called when exiting the resourceType production.
	ExitResourceType(c *ResourceTypeContext)

	// ExitCreateFunctionStatement is called when exiting the createFunctionStatement production.
	ExitCreateFunctionStatement(c *CreateFunctionStatementContext)

	// ExitDropFunctionStatement is called when exiting the dropFunctionStatement production.
	ExitDropFunctionStatement(c *DropFunctionStatementContext)

	// ExitReloadFunctionStatement is called when exiting the reloadFunctionStatement production.
	ExitReloadFunctionStatement(c *ReloadFunctionStatementContext)

	// ExitCreateMacroStatement is called when exiting the createMacroStatement production.
	ExitCreateMacroStatement(c *CreateMacroStatementContext)

	// ExitDropMacroStatement is called when exiting the dropMacroStatement production.
	ExitDropMacroStatement(c *DropMacroStatementContext)

	// ExitCreateViewStatement is called when exiting the createViewStatement production.
	ExitCreateViewStatement(c *CreateViewStatementContext)

	// ExitCreateMaterializedViewStatement is called when exiting the createMaterializedViewStatement production.
	ExitCreateMaterializedViewStatement(c *CreateMaterializedViewStatementContext)

	// ExitViewPartition is called when exiting the viewPartition production.
	ExitViewPartition(c *ViewPartitionContext)

	// ExitDropViewStatement is called when exiting the dropViewStatement production.
	ExitDropViewStatement(c *DropViewStatementContext)

	// ExitDropMaterializedViewStatement is called when exiting the dropMaterializedViewStatement production.
	ExitDropMaterializedViewStatement(c *DropMaterializedViewStatementContext)

	// ExitShowFunctionIdentifier is called when exiting the showFunctionIdentifier production.
	ExitShowFunctionIdentifier(c *ShowFunctionIdentifierContext)

	// ExitShowStmtIdentifier is called when exiting the showStmtIdentifier production.
	ExitShowStmtIdentifier(c *ShowStmtIdentifierContext)

	// ExitTableComment is called when exiting the tableComment production.
	ExitTableComment(c *TableCommentContext)

	// ExitTablePartition is called when exiting the tablePartition production.
	ExitTablePartition(c *TablePartitionContext)

	// ExitTableBuckets is called when exiting the tableBuckets production.
	ExitTableBuckets(c *TableBucketsContext)

	// ExitTableSkewed is called when exiting the tableSkewed production.
	ExitTableSkewed(c *TableSkewedContext)

	// ExitRowFormat is called when exiting the rowFormat production.
	ExitRowFormat(c *RowFormatContext)

	// ExitRecordReader is called when exiting the recordReader production.
	ExitRecordReader(c *RecordReaderContext)

	// ExitRecordWriter is called when exiting the recordWriter production.
	ExitRecordWriter(c *RecordWriterContext)

	// ExitRowFormatSerde is called when exiting the rowFormatSerde production.
	ExitRowFormatSerde(c *RowFormatSerdeContext)

	// ExitRowFormatDelimited is called when exiting the rowFormatDelimited production.
	ExitRowFormatDelimited(c *RowFormatDelimitedContext)

	// ExitTableRowFormat is called when exiting the tableRowFormat production.
	ExitTableRowFormat(c *TableRowFormatContext)

	// ExitTablePropertiesPrefixed is called when exiting the tablePropertiesPrefixed production.
	ExitTablePropertiesPrefixed(c *TablePropertiesPrefixedContext)

	// ExitTableProperties is called when exiting the tableProperties production.
	ExitTableProperties(c *TablePropertiesContext)

	// ExitTablePropertiesList is called when exiting the tablePropertiesList production.
	ExitTablePropertiesList(c *TablePropertiesListContext)

	// ExitKeyValueProperty is called when exiting the keyValueProperty production.
	ExitKeyValueProperty(c *KeyValuePropertyContext)

	// ExitKeyProperty is called when exiting the keyProperty production.
	ExitKeyProperty(c *KeyPropertyContext)

	// ExitTableRowFormatFieldIdentifier is called when exiting the tableRowFormatFieldIdentifier production.
	ExitTableRowFormatFieldIdentifier(c *TableRowFormatFieldIdentifierContext)

	// ExitTableRowFormatCollItemsIdentifier is called when exiting the tableRowFormatCollItemsIdentifier production.
	ExitTableRowFormatCollItemsIdentifier(c *TableRowFormatCollItemsIdentifierContext)

	// ExitTableRowFormatMapKeysIdentifier is called when exiting the tableRowFormatMapKeysIdentifier production.
	ExitTableRowFormatMapKeysIdentifier(c *TableRowFormatMapKeysIdentifierContext)

	// ExitTableRowFormatLinesIdentifier is called when exiting the tableRowFormatLinesIdentifier production.
	ExitTableRowFormatLinesIdentifier(c *TableRowFormatLinesIdentifierContext)

	// ExitTableRowNullFormat is called when exiting the tableRowNullFormat production.
	ExitTableRowNullFormat(c *TableRowNullFormatContext)

	// ExitTableFileFormat is called when exiting the tableFileFormat production.
	ExitTableFileFormat(c *TableFileFormatContext)

	// ExitTableLocation is called when exiting the tableLocation production.
	ExitTableLocation(c *TableLocationContext)

	// ExitColumnNameTypeList is called when exiting the columnNameTypeList production.
	ExitColumnNameTypeList(c *ColumnNameTypeListContext)

	// ExitColumnNameTypeOrConstraintList is called when exiting the columnNameTypeOrConstraintList production.
	ExitColumnNameTypeOrConstraintList(c *ColumnNameTypeOrConstraintListContext)

	// ExitColumnNameColonTypeList is called when exiting the columnNameColonTypeList production.
	ExitColumnNameColonTypeList(c *ColumnNameColonTypeListContext)

	// ExitColumnNameList is called when exiting the columnNameList production.
	ExitColumnNameList(c *ColumnNameListContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitExtColumnName is called when exiting the extColumnName production.
	ExitExtColumnName(c *ExtColumnNameContext)

	// ExitColumnNameOrderList is called when exiting the columnNameOrderList production.
	ExitColumnNameOrderList(c *ColumnNameOrderListContext)

	// ExitColumnParenthesesList is called when exiting the columnParenthesesList production.
	ExitColumnParenthesesList(c *ColumnParenthesesListContext)

	// ExitEnableValidateSpecification is called when exiting the enableValidateSpecification production.
	ExitEnableValidateSpecification(c *EnableValidateSpecificationContext)

	// ExitEnableSpecification is called when exiting the enableSpecification production.
	ExitEnableSpecification(c *EnableSpecificationContext)

	// ExitValidateSpecification is called when exiting the validateSpecification production.
	ExitValidateSpecification(c *ValidateSpecificationContext)

	// ExitEnforcedSpecification is called when exiting the enforcedSpecification production.
	ExitEnforcedSpecification(c *EnforcedSpecificationContext)

	// ExitRelySpecification is called when exiting the relySpecification production.
	ExitRelySpecification(c *RelySpecificationContext)

	// ExitCreateConstraint is called when exiting the createConstraint production.
	ExitCreateConstraint(c *CreateConstraintContext)

	// ExitAlterConstraintWithName is called when exiting the alterConstraintWithName production.
	ExitAlterConstraintWithName(c *AlterConstraintWithNameContext)

	// ExitTableLevelConstraint is called when exiting the tableLevelConstraint production.
	ExitTableLevelConstraint(c *TableLevelConstraintContext)

	// ExitPkUkConstraint is called when exiting the pkUkConstraint production.
	ExitPkUkConstraint(c *PkUkConstraintContext)

	// ExitCheckConstraint is called when exiting the checkConstraint production.
	ExitCheckConstraint(c *CheckConstraintContext)

	// ExitCreateForeignKey is called when exiting the createForeignKey production.
	ExitCreateForeignKey(c *CreateForeignKeyContext)

	// ExitAlterForeignKeyWithName is called when exiting the alterForeignKeyWithName production.
	ExitAlterForeignKeyWithName(c *AlterForeignKeyWithNameContext)

	// ExitSkewedValueElement is called when exiting the skewedValueElement production.
	ExitSkewedValueElement(c *SkewedValueElementContext)

	// ExitSkewedColumnValuePairList is called when exiting the skewedColumnValuePairList production.
	ExitSkewedColumnValuePairList(c *SkewedColumnValuePairListContext)

	// ExitSkewedColumnValuePair is called when exiting the skewedColumnValuePair production.
	ExitSkewedColumnValuePair(c *SkewedColumnValuePairContext)

	// ExitSkewedColumnValues is called when exiting the skewedColumnValues production.
	ExitSkewedColumnValues(c *SkewedColumnValuesContext)

	// ExitSkewedColumnValue is called when exiting the skewedColumnValue production.
	ExitSkewedColumnValue(c *SkewedColumnValueContext)

	// ExitSkewedValueLocationElement is called when exiting the skewedValueLocationElement production.
	ExitSkewedValueLocationElement(c *SkewedValueLocationElementContext)

	// ExitOrderSpecification is called when exiting the orderSpecification production.
	ExitOrderSpecification(c *OrderSpecificationContext)

	// ExitNullOrdering is called when exiting the nullOrdering production.
	ExitNullOrdering(c *NullOrderingContext)

	// ExitColumnNameOrder is called when exiting the columnNameOrder production.
	ExitColumnNameOrder(c *ColumnNameOrderContext)

	// ExitColumnNameCommentList is called when exiting the columnNameCommentList production.
	ExitColumnNameCommentList(c *ColumnNameCommentListContext)

	// ExitColumnNameComment is called when exiting the columnNameComment production.
	ExitColumnNameComment(c *ColumnNameCommentContext)

	// ExitColumnRefOrder is called when exiting the columnRefOrder production.
	ExitColumnRefOrder(c *ColumnRefOrderContext)

	// ExitColumnNameType is called when exiting the columnNameType production.
	ExitColumnNameType(c *ColumnNameTypeContext)

	// ExitColumnNameTypeOrConstraint is called when exiting the columnNameTypeOrConstraint production.
	ExitColumnNameTypeOrConstraint(c *ColumnNameTypeOrConstraintContext)

	// ExitTableConstraint is called when exiting the tableConstraint production.
	ExitTableConstraint(c *TableConstraintContext)

	// ExitColumnNameTypeConstraint is called when exiting the columnNameTypeConstraint production.
	ExitColumnNameTypeConstraint(c *ColumnNameTypeConstraintContext)

	// ExitColumnConstraint is called when exiting the columnConstraint production.
	ExitColumnConstraint(c *ColumnConstraintContext)

	// ExitForeignKeyConstraint is called when exiting the foreignKeyConstraint production.
	ExitForeignKeyConstraint(c *ForeignKeyConstraintContext)

	// ExitColConstraint is called when exiting the colConstraint production.
	ExitColConstraint(c *ColConstraintContext)

	// ExitAlterColumnConstraint is called when exiting the alterColumnConstraint production.
	ExitAlterColumnConstraint(c *AlterColumnConstraintContext)

	// ExitAlterForeignKeyConstraint is called when exiting the alterForeignKeyConstraint production.
	ExitAlterForeignKeyConstraint(c *AlterForeignKeyConstraintContext)

	// ExitAlterColConstraint is called when exiting the alterColConstraint production.
	ExitAlterColConstraint(c *AlterColConstraintContext)

	// ExitColumnConstraintType is called when exiting the columnConstraintType production.
	ExitColumnConstraintType(c *ColumnConstraintTypeContext)

	// ExitDefaultVal is called when exiting the defaultVal production.
	ExitDefaultVal(c *DefaultValContext)

	// ExitTableConstraintType is called when exiting the tableConstraintType production.
	ExitTableConstraintType(c *TableConstraintTypeContext)

	// ExitConstraintOptsCreate is called when exiting the constraintOptsCreate production.
	ExitConstraintOptsCreate(c *ConstraintOptsCreateContext)

	// ExitConstraintOptsAlter is called when exiting the constraintOptsAlter production.
	ExitConstraintOptsAlter(c *ConstraintOptsAlterContext)

	// ExitColumnNameColonType is called when exiting the columnNameColonType production.
	ExitColumnNameColonType(c *ColumnNameColonTypeContext)

	// ExitColType is called when exiting the colType production.
	ExitColType(c *ColTypeContext)

	// ExitColTypeList is called when exiting the colTypeList production.
	ExitColTypeList(c *ColTypeListContext)

	// ExitType_db_col is called when exiting the type_db_col production.
	ExitType_db_col(c *Type_db_colContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitListType is called when exiting the listType production.
	ExitListType(c *ListTypeContext)

	// ExitStructType is called when exiting the structType production.
	ExitStructType(c *StructTypeContext)

	// ExitMapType is called when exiting the mapType production.
	ExitMapType(c *MapTypeContext)

	// ExitUnionType is called when exiting the unionType production.
	ExitUnionType(c *UnionTypeContext)

	// ExitSetOperator is called when exiting the setOperator production.
	ExitSetOperator(c *SetOperatorContext)

	// ExitQueryStatementExpression is called when exiting the queryStatementExpression production.
	ExitQueryStatementExpression(c *QueryStatementExpressionContext)

	// ExitQueryStatementExpressionBody is called when exiting the queryStatementExpressionBody production.
	ExitQueryStatementExpressionBody(c *QueryStatementExpressionBodyContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitCteStatement is called when exiting the cteStatement production.
	ExitCteStatement(c *CteStatementContext)

	// ExitFromStatement is called when exiting the fromStatement production.
	ExitFromStatement(c *FromStatementContext)

	// ExitSingleFromStatement is called when exiting the singleFromStatement production.
	ExitSingleFromStatement(c *SingleFromStatementContext)

	// ExitRegularBody is called when exiting the regularBody production.
	ExitRegularBody(c *RegularBodyContext)

	// ExitAtomSelectStatement is called when exiting the atomSelectStatement production.
	ExitAtomSelectStatement(c *AtomSelectStatementContext)

	// ExitSelectStatement is called when exiting the selectStatement production.
	ExitSelectStatement(c *SelectStatementContext)

	// ExitSetOpSelectStatement is called when exiting the setOpSelectStatement production.
	ExitSetOpSelectStatement(c *SetOpSelectStatementContext)

	// ExitSelectStatementWithCTE is called when exiting the selectStatementWithCTE production.
	ExitSelectStatementWithCTE(c *SelectStatementWithCTEContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitInsertClause is called when exiting the insertClause production.
	ExitInsertClause(c *InsertClauseContext)

	// ExitDestination is called when exiting the destination production.
	ExitDestination(c *DestinationContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitColumnAssignmentClause is called when exiting the columnAssignmentClause production.
	ExitColumnAssignmentClause(c *ColumnAssignmentClauseContext)

	// ExitSetColumnsClause is called when exiting the setColumnsClause production.
	ExitSetColumnsClause(c *SetColumnsClauseContext)

	// ExitUpdateStatement is called when exiting the updateStatement production.
	ExitUpdateStatement(c *UpdateStatementContext)

	// ExitSqlTransactionStatement is called when exiting the sqlTransactionStatement production.
	ExitSqlTransactionStatement(c *SqlTransactionStatementContext)

	// ExitStartTransactionStatement is called when exiting the startTransactionStatement production.
	ExitStartTransactionStatement(c *StartTransactionStatementContext)

	// ExitTransactionMode is called when exiting the transactionMode production.
	ExitTransactionMode(c *TransactionModeContext)

	// ExitTransactionAccessMode is called when exiting the transactionAccessMode production.
	ExitTransactionAccessMode(c *TransactionAccessModeContext)

	// ExitIsolationLevel is called when exiting the isolationLevel production.
	ExitIsolationLevel(c *IsolationLevelContext)

	// ExitLevelOfIsolation is called when exiting the levelOfIsolation production.
	ExitLevelOfIsolation(c *LevelOfIsolationContext)

	// ExitCommitStatement is called when exiting the commitStatement production.
	ExitCommitStatement(c *CommitStatementContext)

	// ExitRollbackStatement is called when exiting the rollbackStatement production.
	ExitRollbackStatement(c *RollbackStatementContext)

	// ExitSetAutoCommitStatement is called when exiting the setAutoCommitStatement production.
	ExitSetAutoCommitStatement(c *SetAutoCommitStatementContext)

	// ExitAbortTransactionStatement is called when exiting the abortTransactionStatement production.
	ExitAbortTransactionStatement(c *AbortTransactionStatementContext)

	// ExitMergeStatement is called when exiting the mergeStatement production.
	ExitMergeStatement(c *MergeStatementContext)

	// ExitWhenClauses is called when exiting the whenClauses production.
	ExitWhenClauses(c *WhenClausesContext)

	// ExitWhenNotMatchedClause is called when exiting the whenNotMatchedClause production.
	ExitWhenNotMatchedClause(c *WhenNotMatchedClauseContext)

	// ExitWhenMatchedAndClause is called when exiting the whenMatchedAndClause production.
	ExitWhenMatchedAndClause(c *WhenMatchedAndClauseContext)

	// ExitWhenMatchedThenClause is called when exiting the whenMatchedThenClause production.
	ExitWhenMatchedThenClause(c *WhenMatchedThenClauseContext)

	// ExitUpdateOrDelete is called when exiting the updateOrDelete production.
	ExitUpdateOrDelete(c *UpdateOrDeleteContext)

	// ExitKillQueryStatement is called when exiting the killQueryStatement production.
	ExitKillQueryStatement(c *KillQueryStatementContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitSelectList is called when exiting the selectList production.
	ExitSelectList(c *SelectListContext)

	// ExitSelectTrfmClause is called when exiting the selectTrfmClause production.
	ExitSelectTrfmClause(c *SelectTrfmClauseContext)

	// ExitSelectItem is called when exiting the selectItem production.
	ExitSelectItem(c *SelectItemContext)

	// ExitTrfmClause is called when exiting the trfmClause production.
	ExitTrfmClause(c *TrfmClauseContext)

	// ExitSelectExpression is called when exiting the selectExpression production.
	ExitSelectExpression(c *SelectExpressionContext)

	// ExitSelectExpressionList is called when exiting the selectExpressionList production.
	ExitSelectExpressionList(c *SelectExpressionListContext)

	// ExitWindow_clause is called when exiting the window_clause production.
	ExitWindow_clause(c *Window_clauseContext)

	// ExitWindow_defn is called when exiting the window_defn production.
	ExitWindow_defn(c *Window_defnContext)

	// ExitWindow_specification is called when exiting the window_specification production.
	ExitWindow_specification(c *Window_specificationContext)

	// ExitWindow_frame is called when exiting the window_frame production.
	ExitWindow_frame(c *Window_frameContext)

	// ExitWindow_range_expression is called when exiting the window_range_expression production.
	ExitWindow_range_expression(c *Window_range_expressionContext)

	// ExitWindow_value_expression is called when exiting the window_value_expression production.
	ExitWindow_value_expression(c *Window_value_expressionContext)

	// ExitWindow_frame_start_boundary is called when exiting the window_frame_start_boundary production.
	ExitWindow_frame_start_boundary(c *Window_frame_start_boundaryContext)

	// ExitWindow_frame_boundary is called when exiting the window_frame_boundary production.
	ExitWindow_frame_boundary(c *Window_frame_boundaryContext)

	// ExitTableAllColumns is called when exiting the tableAllColumns production.
	ExitTableAllColumns(c *TableAllColumnsContext)

	// ExitTableOrColumn is called when exiting the tableOrColumn production.
	ExitTableOrColumn(c *TableOrColumnContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitAliasList is called when exiting the aliasList production.
	ExitAliasList(c *AliasListContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitFromSource is called when exiting the fromSource production.
	ExitFromSource(c *FromSourceContext)

	// ExitAtomjoinSource is called when exiting the atomjoinSource production.
	ExitAtomjoinSource(c *AtomjoinSourceContext)

	// ExitJoinSource is called when exiting the joinSource production.
	ExitJoinSource(c *JoinSourceContext)

	// ExitJoinSourcePart is called when exiting the joinSourcePart production.
	ExitJoinSourcePart(c *JoinSourcePartContext)

	// ExitUniqueJoinSource is called when exiting the uniqueJoinSource production.
	ExitUniqueJoinSource(c *UniqueJoinSourceContext)

	// ExitUniqueJoinExpr is called when exiting the uniqueJoinExpr production.
	ExitUniqueJoinExpr(c *UniqueJoinExprContext)

	// ExitUniqueJoinToken is called when exiting the uniqueJoinToken production.
	ExitUniqueJoinToken(c *UniqueJoinTokenContext)

	// ExitJoinToken is called when exiting the joinToken production.
	ExitJoinToken(c *JoinTokenContext)

	// ExitLateralView is called when exiting the lateralView production.
	ExitLateralView(c *LateralViewContext)

	// ExitTableAlias is called when exiting the tableAlias production.
	ExitTableAlias(c *TableAliasContext)

	// ExitTableBucketSample is called when exiting the tableBucketSample production.
	ExitTableBucketSample(c *TableBucketSampleContext)

	// ExitSplitSample is called when exiting the splitSample production.
	ExitSplitSample(c *SplitSampleContext)

	// ExitTableSample is called when exiting the tableSample production.
	ExitTableSample(c *TableSampleContext)

	// ExitTableSource is called when exiting the tableSource production.
	ExitTableSource(c *TableSourceContext)

	// ExitUniqueJoinTableSource is called when exiting the uniqueJoinTableSource production.
	ExitUniqueJoinTableSource(c *UniqueJoinTableSourceContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitViewName is called when exiting the viewName production.
	ExitViewName(c *ViewNameContext)

	// ExitSubQuerySource is called when exiting the subQuerySource production.
	ExitSubQuerySource(c *SubQuerySourceContext)

	// ExitPartitioningSpec is called when exiting the partitioningSpec production.
	ExitPartitioningSpec(c *PartitioningSpecContext)

	// ExitPartitionTableFunctionSource is called when exiting the partitionTableFunctionSource production.
	ExitPartitionTableFunctionSource(c *PartitionTableFunctionSourceContext)

	// ExitPartitionedTableFunction is called when exiting the partitionedTableFunction production.
	ExitPartitionedTableFunction(c *PartitionedTableFunctionContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitSearchCondition is called when exiting the searchCondition production.
	ExitSearchCondition(c *SearchConditionContext)

	// ExitValuesClause is called when exiting the valuesClause production.
	ExitValuesClause(c *ValuesClauseContext)

	// ExitValuesTableConstructor is called when exiting the valuesTableConstructor production.
	ExitValuesTableConstructor(c *ValuesTableConstructorContext)

	// ExitValueRowConstructor is called when exiting the valueRowConstructor production.
	ExitValueRowConstructor(c *ValueRowConstructorContext)

	// ExitVirtualTableSource is called when exiting the virtualTableSource production.
	ExitVirtualTableSource(c *VirtualTableSourceContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitGroupby_expression is called when exiting the groupby_expression production.
	ExitGroupby_expression(c *Groupby_expressionContext)

	// ExitGroupByEmpty is called when exiting the groupByEmpty production.
	ExitGroupByEmpty(c *GroupByEmptyContext)

	// ExitRollupStandard is called when exiting the rollupStandard production.
	ExitRollupStandard(c *RollupStandardContext)

	// ExitRollupOldSyntax is called when exiting the rollupOldSyntax production.
	ExitRollupOldSyntax(c *RollupOldSyntaxContext)

	// ExitGroupingSetExpression is called when exiting the groupingSetExpression production.
	ExitGroupingSetExpression(c *GroupingSetExpressionContext)

	// ExitGroupingSetExpressionMultiple is called when exiting the groupingSetExpressionMultiple production.
	ExitGroupingSetExpressionMultiple(c *GroupingSetExpressionMultipleContext)

	// ExitGroupingExpressionSingle is called when exiting the groupingExpressionSingle production.
	ExitGroupingExpressionSingle(c *GroupingExpressionSingleContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitHavingCondition is called when exiting the havingCondition production.
	ExitHavingCondition(c *HavingConditionContext)

	// ExitExpressionsInParenthesis is called when exiting the expressionsInParenthesis production.
	ExitExpressionsInParenthesis(c *ExpressionsInParenthesisContext)

	// ExitExpressionsNotInParenthesis is called when exiting the expressionsNotInParenthesis production.
	ExitExpressionsNotInParenthesis(c *ExpressionsNotInParenthesisContext)

	// ExitExpressionPart is called when exiting the expressionPart production.
	ExitExpressionPart(c *ExpressionPartContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitColumnRefOrderInParenthesis is called when exiting the columnRefOrderInParenthesis production.
	ExitColumnRefOrderInParenthesis(c *ColumnRefOrderInParenthesisContext)

	// ExitColumnRefOrderNotInParenthesis is called when exiting the columnRefOrderNotInParenthesis production.
	ExitColumnRefOrderNotInParenthesis(c *ColumnRefOrderNotInParenthesisContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitClusterByClause is called when exiting the clusterByClause production.
	ExitClusterByClause(c *ClusterByClauseContext)

	// ExitPartitionByClause is called when exiting the partitionByClause production.
	ExitPartitionByClause(c *PartitionByClauseContext)

	// ExitDistributeByClause is called when exiting the distributeByClause production.
	ExitDistributeByClause(c *DistributeByClauseContext)

	// ExitSortByClause is called when exiting the sortByClause production.
	ExitSortByClause(c *SortByClauseContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitCastExpression is called when exiting the castExpression production.
	ExitCastExpression(c *CastExpressionContext)

	// ExitCaseExpression is called when exiting the caseExpression production.
	ExitCaseExpression(c *CaseExpressionContext)

	// ExitWhenExpression is called when exiting the whenExpression production.
	ExitWhenExpression(c *WhenExpressionContext)

	// ExitFloorExpression is called when exiting the floorExpression production.
	ExitFloorExpression(c *FloorExpressionContext)

	// ExitFloorDateQualifiers is called when exiting the floorDateQualifiers production.
	ExitFloorDateQualifiers(c *FloorDateQualifiersContext)

	// ExitExtractExpression is called when exiting the extractExpression production.
	ExitExtractExpression(c *ExtractExpressionContext)

	// ExitTimeQualifiers is called when exiting the timeQualifiers production.
	ExitTimeQualifiers(c *TimeQualifiersContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitStringLiteralSequence is called when exiting the stringLiteralSequence production.
	ExitStringLiteralSequence(c *StringLiteralSequenceContext)

	// ExitCharSetStringLiteral is called when exiting the charSetStringLiteral production.
	ExitCharSetStringLiteral(c *CharSetStringLiteralContext)

	// ExitDateLiteral is called when exiting the dateLiteral production.
	ExitDateLiteral(c *DateLiteralContext)

	// ExitTimestampLiteral is called when exiting the timestampLiteral production.
	ExitTimestampLiteral(c *TimestampLiteralContext)

	// ExitTimestampLocalTZLiteral is called when exiting the timestampLocalTZLiteral production.
	ExitTimestampLocalTZLiteral(c *TimestampLocalTZLiteralContext)

	// ExitIntervalValue is called when exiting the intervalValue production.
	ExitIntervalValue(c *IntervalValueContext)

	// ExitIntervalLiteral is called when exiting the intervalLiteral production.
	ExitIntervalLiteral(c *IntervalLiteralContext)

	// ExitIntervalExpression is called when exiting the intervalExpression production.
	ExitIntervalExpression(c *IntervalExpressionContext)

	// ExitIntervalQualifiers is called when exiting the intervalQualifiers production.
	ExitIntervalQualifiers(c *IntervalQualifiersContext)

	// ExitAtomExpression is called when exiting the atomExpression production.
	ExitAtomExpression(c *AtomExpressionContext)

	// ExitPrecedenceUnaryOperator is called when exiting the precedenceUnaryOperator production.
	ExitPrecedenceUnaryOperator(c *PrecedenceUnaryOperatorContext)

	// ExitIsCondition is called when exiting the isCondition production.
	ExitIsCondition(c *IsConditionContext)

	// ExitPrecedenceBitwiseXorOperator is called when exiting the precedenceBitwiseXorOperator production.
	ExitPrecedenceBitwiseXorOperator(c *PrecedenceBitwiseXorOperatorContext)

	// ExitPrecedenceStarOperator is called when exiting the precedenceStarOperator production.
	ExitPrecedenceStarOperator(c *PrecedenceStarOperatorContext)

	// ExitPrecedencePlusOperator is called when exiting the precedencePlusOperator production.
	ExitPrecedencePlusOperator(c *PrecedencePlusOperatorContext)

	// ExitPrecedenceConcatenateOperator is called when exiting the precedenceConcatenateOperator production.
	ExitPrecedenceConcatenateOperator(c *PrecedenceConcatenateOperatorContext)

	// ExitPrecedenceAmpersandOperator is called when exiting the precedenceAmpersandOperator production.
	ExitPrecedenceAmpersandOperator(c *PrecedenceAmpersandOperatorContext)

	// ExitPrecedenceBitwiseOrOperator is called when exiting the precedenceBitwiseOrOperator production.
	ExitPrecedenceBitwiseOrOperator(c *PrecedenceBitwiseOrOperatorContext)

	// ExitPrecedenceRegexpOperator is called when exiting the precedenceRegexpOperator production.
	ExitPrecedenceRegexpOperator(c *PrecedenceRegexpOperatorContext)

	// ExitPrecedenceSimilarOperator is called when exiting the precedenceSimilarOperator production.
	ExitPrecedenceSimilarOperator(c *PrecedenceSimilarOperatorContext)

	// ExitPrecedenceDistinctOperator is called when exiting the precedenceDistinctOperator production.
	ExitPrecedenceDistinctOperator(c *PrecedenceDistinctOperatorContext)

	// ExitPrecedenceEqualOperator is called when exiting the precedenceEqualOperator production.
	ExitPrecedenceEqualOperator(c *PrecedenceEqualOperatorContext)

	// ExitPrecedenceNotOperator is called when exiting the precedenceNotOperator production.
	ExitPrecedenceNotOperator(c *PrecedenceNotOperatorContext)

	// ExitPrecedenceAndOperator is called when exiting the precedenceAndOperator production.
	ExitPrecedenceAndOperator(c *PrecedenceAndOperatorContext)

	// ExitPrecedenceOrOperator is called when exiting the precedenceOrOperator production.
	ExitPrecedenceOrOperator(c *PrecedenceOrOperatorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSubQueryExpression is called when exiting the subQueryExpression production.
	ExitSubQueryExpression(c *SubQueryExpressionContext)

	// ExitPrecedenceSimilarExpressionPart is called when exiting the precedenceSimilarExpressionPart production.
	ExitPrecedenceSimilarExpressionPart(c *PrecedenceSimilarExpressionPartContext)

	// ExitPrecedenceSimilarExpressionAtom is called when exiting the precedenceSimilarExpressionAtom production.
	ExitPrecedenceSimilarExpressionAtom(c *PrecedenceSimilarExpressionAtomContext)

	// ExitPrecedenceSimilarExpressionIn is called when exiting the precedenceSimilarExpressionIn production.
	ExitPrecedenceSimilarExpressionIn(c *PrecedenceSimilarExpressionInContext)

	// ExitPrecedenceSimilarExpressionPartNot is called when exiting the precedenceSimilarExpressionPartNot production.
	ExitPrecedenceSimilarExpressionPartNot(c *PrecedenceSimilarExpressionPartNotContext)

	// ExitBooleanValue is called when exiting the booleanValue production.
	ExitBooleanValue(c *BooleanValueContext)

	// ExitBooleanValueTok is called when exiting the booleanValueTok production.
	ExitBooleanValueTok(c *BooleanValueTokContext)

	// ExitTableOrPartition is called when exiting the tableOrPartition production.
	ExitTableOrPartition(c *TableOrPartitionContext)

	// ExitPartitionSpec is called when exiting the partitionSpec production.
	ExitPartitionSpec(c *PartitionSpecContext)

	// ExitPartitionVal is called when exiting the partitionVal production.
	ExitPartitionVal(c *PartitionValContext)

	// ExitDropPartitionSpec is called when exiting the dropPartitionSpec production.
	ExitDropPartitionSpec(c *DropPartitionSpecContext)

	// ExitDropPartitionVal is called when exiting the dropPartitionVal production.
	ExitDropPartitionVal(c *DropPartitionValContext)

	// ExitDropPartitionOperator is called when exiting the dropPartitionOperator production.
	ExitDropPartitionOperator(c *DropPartitionOperatorContext)

	// ExitSysFuncNames is called when exiting the sysFuncNames production.
	ExitSysFuncNames(c *SysFuncNamesContext)

	// ExitDescFuncNames is called when exiting the descFuncNames production.
	ExitDescFuncNames(c *DescFuncNamesContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitFunctionIdentifier is called when exiting the functionIdentifier production.
	ExitFunctionIdentifier(c *FunctionIdentifierContext)

	// ExitPrincipalIdentifier is called when exiting the principalIdentifier production.
	ExitPrincipalIdentifier(c *PrincipalIdentifierContext)

	// ExitNonReserved is called when exiting the nonReserved production.
	ExitNonReserved(c *NonReservedContext)

	// ExitSql11ReservedKeywordsUsedAsFunctionName is called when exiting the sql11ReservedKeywordsUsedAsFunctionName production.
	ExitSql11ReservedKeywordsUsedAsFunctionName(c *Sql11ReservedKeywordsUsedAsFunctionNameContext)

	// ExitResourcePlanDdlStatements is called when exiting the resourcePlanDdlStatements production.
	ExitResourcePlanDdlStatements(c *ResourcePlanDdlStatementsContext)

	// ExitRpAssign is called when exiting the rpAssign production.
	ExitRpAssign(c *RpAssignContext)

	// ExitRpAssignList is called when exiting the rpAssignList production.
	ExitRpAssignList(c *RpAssignListContext)

	// ExitRpUnassign is called when exiting the rpUnassign production.
	ExitRpUnassign(c *RpUnassignContext)

	// ExitRpUnassignList is called when exiting the rpUnassignList production.
	ExitRpUnassignList(c *RpUnassignListContext)

	// ExitCreateResourcePlanStatement is called when exiting the createResourcePlanStatement production.
	ExitCreateResourcePlanStatement(c *CreateResourcePlanStatementContext)

	// ExitWithReplace is called when exiting the withReplace production.
	ExitWithReplace(c *WithReplaceContext)

	// ExitActivate is called when exiting the activate production.
	ExitActivate(c *ActivateContext)

	// ExitEnable is called when exiting the enable production.
	ExitEnable(c *EnableContext)

	// ExitDisable is called when exiting the disable production.
	ExitDisable(c *DisableContext)

	// ExitUnmanaged is called when exiting the unmanaged production.
	ExitUnmanaged(c *UnmanagedContext)

	// ExitAlterResourcePlanStatement is called when exiting the alterResourcePlanStatement production.
	ExitAlterResourcePlanStatement(c *AlterResourcePlanStatementContext)

	// ExitGlobalWmStatement is called when exiting the globalWmStatement production.
	ExitGlobalWmStatement(c *GlobalWmStatementContext)

	// ExitReplaceResourcePlanStatement is called when exiting the replaceResourcePlanStatement production.
	ExitReplaceResourcePlanStatement(c *ReplaceResourcePlanStatementContext)

	// ExitDropResourcePlanStatement is called when exiting the dropResourcePlanStatement production.
	ExitDropResourcePlanStatement(c *DropResourcePlanStatementContext)

	// ExitPoolPath is called when exiting the poolPath production.
	ExitPoolPath(c *PoolPathContext)

	// ExitTriggerExpression is called when exiting the triggerExpression production.
	ExitTriggerExpression(c *TriggerExpressionContext)

	// ExitTriggerExpressionStandalone is called when exiting the triggerExpressionStandalone production.
	ExitTriggerExpressionStandalone(c *TriggerExpressionStandaloneContext)

	// ExitTriggerOrExpression is called when exiting the triggerOrExpression production.
	ExitTriggerOrExpression(c *TriggerOrExpressionContext)

	// ExitTriggerAndExpression is called when exiting the triggerAndExpression production.
	ExitTriggerAndExpression(c *TriggerAndExpressionContext)

	// ExitTriggerAtomExpression is called when exiting the triggerAtomExpression production.
	ExitTriggerAtomExpression(c *TriggerAtomExpressionContext)

	// ExitTriggerLiteral is called when exiting the triggerLiteral production.
	ExitTriggerLiteral(c *TriggerLiteralContext)

	// ExitComparisionOperator is called when exiting the comparisionOperator production.
	ExitComparisionOperator(c *ComparisionOperatorContext)

	// ExitTriggerActionExpression is called when exiting the triggerActionExpression production.
	ExitTriggerActionExpression(c *TriggerActionExpressionContext)

	// ExitTriggerActionExpressionStandalone is called when exiting the triggerActionExpressionStandalone production.
	ExitTriggerActionExpressionStandalone(c *TriggerActionExpressionStandaloneContext)

	// ExitCreateTriggerStatement is called when exiting the createTriggerStatement production.
	ExitCreateTriggerStatement(c *CreateTriggerStatementContext)

	// ExitAlterTriggerStatement is called when exiting the alterTriggerStatement production.
	ExitAlterTriggerStatement(c *AlterTriggerStatementContext)

	// ExitDropTriggerStatement is called when exiting the dropTriggerStatement production.
	ExitDropTriggerStatement(c *DropTriggerStatementContext)

	// ExitPoolAssign is called when exiting the poolAssign production.
	ExitPoolAssign(c *PoolAssignContext)

	// ExitPoolAssignList is called when exiting the poolAssignList production.
	ExitPoolAssignList(c *PoolAssignListContext)

	// ExitCreatePoolStatement is called when exiting the createPoolStatement production.
	ExitCreatePoolStatement(c *CreatePoolStatementContext)

	// ExitAlterPoolStatement is called when exiting the alterPoolStatement production.
	ExitAlterPoolStatement(c *AlterPoolStatementContext)

	// ExitDropPoolStatement is called when exiting the dropPoolStatement production.
	ExitDropPoolStatement(c *DropPoolStatementContext)

	// ExitCreateMappingStatement is called when exiting the createMappingStatement production.
	ExitCreateMappingStatement(c *CreateMappingStatementContext)

	// ExitAlterMappingStatement is called when exiting the alterMappingStatement production.
	ExitAlterMappingStatement(c *AlterMappingStatementContext)

	// ExitDropMappingStatement is called when exiting the dropMappingStatement production.
	ExitDropMappingStatement(c *DropMappingStatementContext)
}
