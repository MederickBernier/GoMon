param (
    [string[]]$Extensions = @("*"), # Liste des extensions (par défaut tous les fichiers)
    [switch]$ExcludeEmptyAndComments # Switch pour exclure les lignes vides et les commentaires
)

# Initialiser une variable pour le total des lignes
$totalLines = 0

# Parcourir toutes les extensions fournies
foreach ($ext in $Extensions) {
    # Parcourir tous les fichiers correspondant à l'extension
    Get-ChildItem -Recurse -Filter "*.$ext" | ForEach-Object {
        # Compter les lignes en fonction du switch
        $lines = if ($ExcludeEmptyAndComments.IsPresent) {
            # Exclure lignes vides et commentaires
            (Get-Content $_.FullName | Where-Object { $_ -notmatch "^\s*$" -and $_ -notmatch "^\s*\/\/" }).Count
        }
        else {
            # Compter toutes les lignes
            (Get-Content $_.FullName).Count
        }
        # Afficher le fichier et le nombre de lignes
        Write-Output "$($_.FullName): $lines lignes"
        # Ajouter au total
        $totalLines += $lines
    }
}

# Afficher le total
Write-Output "`nTotal des lignes de code pour les extensions $($Extensions -join ', '): $totalLines lignes"
